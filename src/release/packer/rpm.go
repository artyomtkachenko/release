package packer

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"release/meta"
)

const (
	rpmSpec = `
# Disable the stupid stuff rpm distros include in the build process by default:
#   Disable any prep shell actions. replace them with simply 'true'
%define __spec_prep_post true
%define __spec_prep_pre true
#   Disable any build shell actions. replace them with simply 'true'
%define __spec_build_post true
%define __spec_build_pre true
#   Disable any install shell actions. replace them with simply 'true'
%define __spec_install_post true
%define __spec_install_pre true
#   Disable any clean shell actions. replace them with simply 'true'
%define __spec_clean_post true
%define __spec_clean_pre true
# Disable checking for unpackaged files ?
#%undefine __check_files

# Use md5 file digest method
%define _binary_filedigest_algorithm 1

# Use gzip payload compression
%define _binary_payload w9.gzdio 

#Define a custom output directory
%define _rpmdir {{ .RpmsDir }}
%define _source_filedigest_algorithm 1

Name: {{ .ReleaseMeta.Project.Name }}
Version: {{ .ReleaseMeta.Project.Version }}        
Release:        1
Summary: {{ .ReleaseMeta.Project.Description }}        
AutoReqProv: no
# Seems specifying BuildRoot is required on older rpmbuild (like on CentOS 5)
# fpm passes '--define buildroot ...' on the commandline, so just reuse that.
BuildRoot: %buildroot
# Add prefix, must not end with /
Prefix: {{ .ReleaseMeta.Deploy.RootDir }}

Group: default
License: Private        
Vendor: {{ .Vendor }} 
URL: {{ .ReleaseMeta.Project.ScmUrl }}            
Packager: {{ .ReleaseMeta.Project.Email }}

%description
{{ .ReleaseMeta.Project.Description }}

%prep
# noop

%build
# noop

%install
# noop

%clean
# noop

%post{{ index .Scripts "after-install.sh" }}

%preun{{ index .Scripts "before-remove.sh" }}

%files
%defattr(-,{{ .ReleaseMeta.Deploy.User }},{{ .ReleaseMeta.Deploy.Group }},-)
{{ $user := .ReleaseMeta.Deploy.User }}
{{ $group := .ReleaseMeta.Deploy.Group }}
{{ range .Dirs }}
%dir %attr({{ .Mode }}, {{ $user }}, {{ $group }}) {{ .Path }}{{end}}
{{ range .Files }}
%attr({{ .Mode }}, {{ $user }}, {{ $group }}) {{ .Path }}{{end}}

%changelog
`
)

type tmpl struct {
	ReleaseMeta meta.ReleaseMeta
	Files       []file
	Dirs        []file
	Scripts     map[string]string
	Vendor      string
	RpmsDir     string
}

type file struct {
	Path string
	Mode string
}

var dirs []file
var files []file

func convertBuildDirToDepoyDir(buildRoot string, appInstallRoot string, files []file) []file {
	var result []file

	for _, f := range files {
		a := file{
			Path: strings.Replace(f.Path, buildRoot, "", -1),
			Mode: f.Mode,
		}
		result = append(result, a)
	}
	return result
}

func walkBuildDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
		return nil
	}

	//This conversion is so hacky. Does convertion from int32 to octal
	octal := strconv.FormatInt(int64(info.Mode()), 8)
	meta := file{
		path,
		"0" + octal[len(octal)-3:len(octal)],
	}

	if info.IsDir() {
		dirs = append(dirs, meta)
	} else {
		files = append(files, meta)
	}
	return nil
}

func getScripts(path string) (map[string]string, error) {
	fileInfo, err := os.Stat(path)
	var result = make(map[string]string)
	if err != nil {
		return result, err
	}
	if fileInfo.IsDir() {
		scripts, err := filepath.Glob(filepath.Join(path, "*.sh"))
		if err != nil {
			return result, err
		}

		for _, script := range scripts {
			fi, err := os.Stat(script)
			if err != nil {
				return result, err
			}
			if !fi.IsDir() {
				fh, err := os.Open(script)
				if err != nil {
					return result, err
				}
				content, err := ioutil.ReadAll(fh)
				if err != nil {
					return result, err
				}
				result[filepath.Base(script)] = string(content)
			}
		}
	}
	return result, nil
}

func GenerateRpmSpec(rm meta.ReleaseMeta, buildRoot string) error {
	t := template.New("RPM SPEC template")
	t, err := t.Parse(rpmSpec)

	specDir := filepath.Join(buildRoot, "SPEC")
	rpmsDir := filepath.Join(buildRoot, "RPMS")
	buildDir := filepath.Join(buildRoot, "BUILD")
	var scripts map[string]string

	if rm.Scripts.BeforeRemove != "" || rm.Scripts.AfterInstall != "" {

		oldScriptsDir := filepath.Join(buildDir, "__SCRIPTS__") //TODO make it more generic, based on the metadata provided
		newScriptsDir := filepath.Join(buildRoot, "__SCRIPTS__")
		if err := os.Rename(oldScriptsDir, newScriptsDir); err != nil {
			return err
		}

		scripts, err = getScripts(newScriptsDir)
		if err != nil {
			return err
		}
	}

	if err := filepath.Walk(buildDir, walkBuildDir); err != nil {
		return err
	}
	specFile := filepath.Join(specDir, rm.Project.Name+".spec")

	dirsToCreate := []string{specDir, rpmsDir}
	for _, dir := range dirsToCreate {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	f, err := os.Create(specFile)
	defer f.Close()
	if err != nil {
		return err
	}
	// files and dirs are local vars to to this package
	vendor, _ := os.Hostname()
	templateData := tmpl{
		rm,
		convertBuildDirToDepoyDir(buildDir, rm.Deploy.RootDir, files),
		convertBuildDirToDepoyDir(buildDir, rm.Deploy.RootDir, dirs),
		scripts,
		vendor,
		rpmsDir,
	}

	err = t.Execute(f, templateData)
	if err != nil {
		return err
	}
	return nil
}

func RunRpmBuild(rm meta.ReleaseMeta, buildRoot string) error {
	buildDir, _ := filepath.Abs(filepath.Join(buildRoot, "BUILD"))
	specFile, _ := filepath.Abs(filepath.Join(buildRoot, "SPEC", rm.Project.Name+".spec"))
	cmd := "rpmbuild --clean  -bb --buildroot " + buildDir + " " + specFile

	_, err := exec.Command("sh", "-c", cmd).Output()
	return err
}
