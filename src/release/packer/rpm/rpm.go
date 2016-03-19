package rpm

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"release/config"
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

Name: {{ .ReleaseConfig.Project.Name }}
Version: {{ .ReleaseConfig.Project.Version }}        
Release:        1
Summary: {{ .ReleaseConfig.Project.Description }}        
AutoReqProv: no
# Seems specifying BuildRoot is required on older rpmbuild (like on CentOS 5)
# fpm passes '--define buildroot ...' on the commandline, so just reuse that.
BuildRoot: %buildroot
# Add prefix, must not end with /
Prefix: {{ .ReleaseConfig.Deploy.RootDir }}

Group: default
License: Private        
Vendor: {{ .Vendor }} 
URL: {{ .ReleaseConfig.Project.ScmUrl }}            
Packager: {{ .ReleaseConfig.Project.Email }}

%description
{{ .ReleaseConfig.Project.Description }}

%prep
# noop

%build
# noop

%install
# noop

%clean
# noop

%post
{{ index .Scripts "post" }}

%preun
{{ index .Scripts "preun" }}

%files
%defattr(-,{{ .ReleaseConfig.Deploy.User }},{{ .ReleaseConfig.Deploy.Group }},-)
{{ $user := .ReleaseConfig.Deploy.User }}
{{ $group := .ReleaseConfig.Deploy.Group }}
{{ range .Dirs }}
%dir %attr({{ .Mode }}, {{ $user }}, {{ $group }}) {{ .Path }}{{end}}
{{ range .Files }}
%attr({{ .Mode }}, {{ $user }}, {{ $group }}) {{ .Path }}{{end}}

%changelog
`
)

type tmpl struct {
	ReleaseConfig config.ReleaseConfig
	Files         []file
	Dirs          []file
	Scripts       map[string]string
	Vendor        string
	RpmsDir       string
}

type file struct {
	Path string
	Mode string
}

var dirs []file
var files []file

func GenerateRpmBuildDirs(buildRoot string, rc config.ReleaseConfig) error {
	buildDirs := []string{"BUILD", "SPEC", "RPMS"}
	for _, bd := range buildDirs {
		dir := filepath.Join(buildRoot, bd)
		if bd == "BUILD" {
			dir = filepath.Join(dir, rc.Deploy.RootDir)
		}
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func getBuildDir(buildRoot string) string {
	return filepath.Join(buildRoot, "BUILD")
}

func getSpecDir(buildRoot string) string {
	return filepath.Join(buildRoot, "SPEC")
}

func getRpmsDir(buildRoot string) string {
	return filepath.Join(buildRoot, "RPMS")
}

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

//Free all global resources
func free() {
	files = nil
	dirs = nil
}

func GenerateRpmSpec(rm config.ReleaseConfig, buildRoot string, scripts map[string]string) error {
	t := template.New("RPM SPEC template")
	t, err := t.Parse(rpmSpec)

	specDir := getSpecDir(buildRoot)
	rpmsDir := getRpmsDir(buildRoot)
	buildDir := getBuildDir(buildRoot)

	if err := filepath.Walk(buildDir, walkBuildDir); err != nil {
		return err
	}
	specFile := filepath.Join(specDir, rm.Project.Name+".spec")

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
	defer free()
	return nil
}

func RunRpmBuild(rc config.ReleaseConfig, buildRoot string) error {
	buildDir, _ := filepath.Abs(getBuildDir(buildRoot))
	specFile, _ := filepath.Abs(filepath.Join(getSpecDir(buildRoot), rc.Project.Name+".spec"))
	sign := ""
	if rc.Package.Sign {
		sign = " --sign "
	}
	cmd := "rpmbuild --clean  -bb --target " +
		rc.Project.Arch +
		" --buildroot " + buildDir +
		sign +
		" " + specFile

	fmt.Println(cmd)
	_, err := exec.Command("sh", "-c", cmd).Output()

	return err
}
