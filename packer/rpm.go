package packer

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/artyomtkachenko/release/config"
	"github.com/artyomtkachenko/release/meta"
)

const (
	rpmSpec = `
Name: {{ .ReleaseMeta.Project.Name }}
Version: {{ .ReleaseMeta.Project.Version }}        
Release:        1%{?dist}
Summary: {{ .ReleaseMeta.Project.Description }}        
Packager: Santa Claus <sclaus@northpole.com>

Group: {{ .ReleaseMeta.Project.Email }}           
License: Private        
URL: {{ .ReleaseMeta.Project.ScmUrl }}            
BuildRoot: {{ .ReleaseMeta.Project.BuildRoot }}

%description
{{ .ReleaseMeta.Project.Description }}

%files
%defattr(-,{{ .ReleaseMeta.Deploy.User }},{{ .ReleaseMeta.Deploy.Group }},-)
{{ $user := .ReleaseMeta.Deploy.User }}
{{ $group := .ReleaseMeta.Deploy.Group }}
{{ range .Dirs }}
%dir %attr({{ .Mode }}, {{ $user }}, {{ $group }}) {{ .Path }}{{end}}
{{ range .Files }}
%attr({{ .Mode }}, {{ $user }}, {{ $group }}) {{ .Path }}{{end}}
`
)

type tmpl struct {
	ReleaseMeta meta.ReleaseMeta
	Files       []file
	Dirs        []file
}
type file struct {
	Path string
	Mode string
}

var dirs []file
var files []file

func printFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
		return nil
	}
	//This conversion is so hacky
	octal := strconv.FormatInt(int64(info.Mode()), 8)
	meta := file{path, "0" + octal[len(octal)-3:len(octal)]}

	if info.IsDir() {
		dirs = append(dirs, meta)
	} else {
		files = append(files, meta)
	}
	return nil
}

//Converts JSON object into RPM spec file
func GenerateRpmSpec(rm meta.ReleaseMeta, conf config.Config, tmp TmpDir) error {
	t := template.New("RPM SPEC template")
	t, err := t.Parse(rpmSpec)

	specDir := filepath.Join(conf.DataDir, tmp.Path, "SPEC")
	buildDir := filepath.Join(conf.DataDir, tmp.Path, "BUILD")
	err = filepath.Walk(buildDir, printFile)
	if err != nil {
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
	templateData := tmpl{rm, files, dirs}
	fmt.Printf("%+v\n", templateData)

	err = t.Execute(f, templateData)
	if err != nil {
		return err
	}
	return nil
}
