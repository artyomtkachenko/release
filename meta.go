package main

type ReleaseDeploy struct {
	User    string `json:"user"`
	Group   string `json:"group"`
	RootDir string `json:"root_dir"`
}

type ReleasePackage struct {
	Sign bool   `json:"root_dir"`
	Type string `json:"package_type"`
}

type ReleasePublish struct {
	Type        string `json:"publish_type"`
	Destination string `json:"destination"`
}

type ReleaseScripts struct {
	BeforeRemove string `json:"before_remove"`
	AfterInstall string `json:"after_install"`
}

type ReleaseProject struct {
	Name        string `json:"name"`
	ContentRoot string `json:"content_root"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type ReleaseMeta struct {
	ReleaseProject
	ReleaseDeploy
	ReleasePackage
	ReleasePublish
	ReleaseScripts
}
