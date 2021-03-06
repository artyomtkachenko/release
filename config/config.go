package config

type ServerConfig struct {
	Port    string
	DataDir string
	Debug   bool
}

type MetaData struct {
	Version  string
	Revision string
}

type Deploy struct {
	User    string `json:"user" yaml:"user"`
	Group   string `json:"group" yaml:"group"`
	RootDir string `json:"root_dir,app_root" yaml:"root_dir,app_root"`
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
	ContentRoot string `json:"content_root" yaml:"content_root"`
	Email       string `json:"email" yaml:"email"`
	Description string `json:"description" yaml:"description"`
	ScmUrl      string `json:"url" yaml:"url"`
	Version     string `json:"version" yaml:"version"`
	Revision    string `json:"revision" yaml:"revision"`
	Arch        string `json:"arch" yaml:"arch"`
}

type ReleaseConfig struct {
	Project `json:"project" yaml:"project"`
	Deploy  `json:"deploy"  yaml:"deploy"`
	Package `json:"package" yaml:"package"`
	Publish `json:"publish" yaml:"publish"`
	Scripts `json:"scripts" yaml:"scripts"`
}
