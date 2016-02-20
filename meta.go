package main

type Deploy struct {
	User    string `json:"user" yaml:"user"`
	Group   string `json:"group" yaml:"group"`
	RootDir string `json:"root_dir" yaml:"root_dir"`
}

type Package struct {
	Sign bool   `json:"sign" yaml:"sign"`
	Type string `json:"type" yaml:"type"`
}

type Publish struct {
	Type        string `json:"type" yaml:"type"`
	Destination string `json:"destination" yaml:"destination"`
}

type Scripts struct {
	BeforeRemove string `json:"before_remove" yaml:"before_remove"`
	AfterInstall string `json:"after_install" yaml:"after_install"`
}

type Project struct {
	Name        string `json:"name" yaml:"name"`
	BuildRoot   string `json:"build_root" yaml:"build_root"`
	Email       string `json:"email" yaml:"email"`
	Description string `json:"description" yaml:"description"`
	ScmUrl      string `json:"url" yaml:"url"`
	Version     string `json:"version" yaml:"version"`
}

type ReleaseMeta struct {
	Project `json:"project" yaml:"project"`
	Deploy  `json:"deploy"  yaml:"deploy"`
	Package `json:"package" yaml:"package"`
	Publish `json:"publish" yaml:"publish"`
	Scripts `json:"scripts" yaml:"scripts"`
}
