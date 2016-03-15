package meta

import (
	"os"
	"path/filepath"
	"release/config"
	"testing"
)

func TestSaveMeta(t *testing.T) {
	conf := config.Config{
		DataDir: "testdata",
	}
	m := ReleaseMeta{
		Project: Project{
			Name:        "foo",
			BuildRoot:   "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Scripts: Scripts{
			BeforeRemove: "before-remove.sh",
			AfterInstall: "after-install.sh",
		},
		Publish: Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: Deploy{
			User:    "bob",
			Group:   "bob",
			RootDir: "/root",
		},
	}

	os.MkdirAll(filepath.Join(conf.DataDir, "bbb"), 0755)
	if err := SaveMeta(m, conf, "bbb"); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("testdata/bbb")
}
