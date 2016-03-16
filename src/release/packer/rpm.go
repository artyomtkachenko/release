package packer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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

Name: {{ .ReleaseMeta.Project.Name }}
Version: {{ .ReleaseMeta.Project.Version }}        
Release:        1
Summary: {{ .ReleaseMeta.Project.Description }}        
AutoReqProv: no
#BuildRoot: {{ .ReleaseMeta.Project.BuildRoot }}
# Seems specifying BuildRoot is required on older rpmbuild (like on CentOS 5)
# fpm passes '--define buildroot ...' on the commandline, so just reuse that.
BuildRoot: %buildroot
# Add prefix, must not end with /
Prefix: {{ .ReleaseMeta.Deploy.RootDir }}

Group: default
License: Private        
Vendor: 
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
}

type file struct {
	Path string
	Mode string
}

var dirs []file
var files []file

func walkBuildDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
		return nil
	}
	//This conversion is so hacky. Does convertion from int32 to octal
	octal := strconv.FormatInt(int64(info.Mode()), 8)
	meta := file{path, "0" + octal[len(octal)-3:len(octal)]}

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
	buildDir := filepath.Join(buildRoot, "BUILD")
	var scripts map[string]string

	if rm.Scripts.BeforeRemove != "" || rm.Scripts.AfterInstall != "" {
		scriptsdDir := filepath.Join(buildDir, "__SCRIPTS__")
		scripts, err = getScripts(scriptsdDir)
		if err != nil {
			return err
		}
	}

	if err := filepath.Walk(buildDir, walkBuildDir); err != nil {
		return err
	}
	specFile := filepath.Join(specDir, rm.Project.Name+".spec")
	if err := os.MkdirAll(specDir, 0755); err != nil {
		return err
	}

	f, err := os.Create(specFile)
	defer f.Close()
	if err != nil {
		return err
	}
	// files and dirs are local vars to to this package
	templateData := tmpl{rm, files, dirs, scripts}

	err = t.Execute(f, templateData)
	if err != nil {
		return err
	}
	return nil
}
