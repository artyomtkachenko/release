package packer

import (
	"github.com/artyomtkachenko/release/config"
	"github.com/artyomtkachenko/release/meta"
	"os"
	"path/filepath"
	"text/template"
)

const (
	rpmSpec = `
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
`
)

//Converts JSON object into RPM spec file
func ConvertJSON2RpmSpec(rm meta.ReleaseMeta, conf config.Config, tmp TmpDir) error {
	t := template.New("RPM SPEC template")
	t, err := t.Parse(rpmSpec)

	specDir := filepath.Join(conf.DataDir, tmp.Path, "SPEC")
	specFile := filepath.Join(specDir, rm.Project.Name)
	if err := os.MkdirAll(specDir, 0755); err != nil {
		return err
	}

	f, err := os.Create(specFile)
	if err != nil {
		return err
	}
	err = t.Execute(f, rm)
	return err
}
