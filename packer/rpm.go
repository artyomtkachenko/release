package packer

import (
	"fmt"
	"github.com/artyomtkachenko/release/config"
	"github.com/artyomtkachenko/release/meta"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
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
{{ range .files }}
%defattr(-,{{ .user }},{{ .group }},-)
{{ end }}
{{ end }}
`
)

type attr struct {
	user  string
	group string
	files *[]os.FileInfo
}

//Converts JSON object into RPM spec file
func GenerateRpmSpec(rm meta.ReleaseMeta, conf config.Config, tmp TmpDir) error {
	t := template.New("RPM SPEC template")
	t, err := t.Parse(rpmSpec)

	specDir := filepath.Join(conf.DataDir, tmp.Path, "SPEC")
	buildDir := filepath.Join(conf.DataDir, tmp.Path, "BUILD")
	files, err := ioutil.ReadDir(buildDir)
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
	attrs := attr{rm.Deploy.User, rm.Deploy.Group, &files}
	for _, ff := range files {
		fmt.Printf("%+v\n", ff.Name())
	}
	err = t.ExecuteTemplate(f, "T1", rm)
	return err
	err = t.ExecuteTemplate(f, "T2", attrs)
	return err
}
