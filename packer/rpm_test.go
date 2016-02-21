package packer

import (
	"github.com/artyomtkachenko/release/config"
	"github.com/artyomtkachenko/release/meta"
	"os"
	"testing"
)

func TestConvertJSON2RpmSpec(t *testing.T) {
	conf := config.Config{
		DataDir: "/tmp",
	}

	m := meta.ReleaseMeta{
		Project: meta.Project{
			Name:        "foo",
			BuildRoot:   "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Publish: meta.Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: meta.Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: meta.Deploy{
			User:    "bob",
			Group:   "bob",
			RootDir: "/root",
		},
	}

	tmpId := TmpDir{Path: "aaa"}
	err := GenerateRpmSpec(m, conf, tmpId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	defer os.RemoveAll("/tmp/aaa")
}