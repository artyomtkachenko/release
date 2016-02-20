package main

type Deploy struct {
	User    string `json:"user"`
	Group   string `json:"group"`
	RootDir string `json:"root_dir"`
}

type Package struct {
	Sign bool   `json:"sign"`
	Type string `json:"package_type"`
}

type Publish struct {
	Type        string `json:"publish_type"`
	Destination string `json:"destination"`
}

type Scripts struct {
	BeforeRemove string `json:"before_remove"`
	AfterInstall string `json:"after_install"`
}

type Project struct {
	Name        string `json:"name"`
	ContentRoot string `json:"content_root"`
	Email       string `json:"email"`
	Description string `json:"description"`
	ScmUrl      string `json:"url"`
}

type ReleaseMeta struct {
	Project `json:"project"`
	Deploy  `json:"deploy"`
	Package `json:"package"`
	Publish `json:"publish"`
	Scripts `json:"scripts"`
}
