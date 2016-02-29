package packer

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/artyomtkachenko/release/config"
	"github.com/artyomtkachenko/release/meta"
)

const (
	rpmSpec = `
{{ define "T1" }}	
Name: {{ .Project.Name }}
Version: {{ .Project.Version }}        
Release:        1%{?dist}
Summary: {{ .Project.Description }}        
Packager: Santa Claus <sclaus@northpole.com>

Group: {{ .Project.Email }}           
License: Private        
URL: {{ .Project.ScmUrl }}            
BuildRoot: {{ .Project.BuildRoot }}

%description
{{ .Project.Description }}

%files
%defattr(-,{{ .Deploy.User }},{{ .Deploy.Group }},-)
{{ end }}
{{ define "T2" }}	
{{ range .Dirs }}
%dir %attr({{ .Mode }}, go, go) {{ .Path }}{{end}}
{{ range .Files }}
%attr({{ .Mode }}, go, go) {{ .Path }}{{end}}
{{ end }}
`
)

type attr struct {
	user  string
	group string
	Files []file
	Dirs  []file
}

type file struct {
	Path string
	Mode os.FileMode
}

var dirs []file
var files []file

func printFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
		return nil
	}
	meta := file{path, info.Mode()}

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
	attrs := attr{rm.Deploy.User, rm.Deploy.Group, files, dirs}
	fmt.Printf("%+v\n", attrs)
	// err = t.ExecuteTemplate(f, "T1", rm)
	// return err
	err = t.ExecuteTemplate(f, "T2", attrs)
	return err
}
