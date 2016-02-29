package packer

import (
	"encoding/json"
	"github.com/artyomtkachenko/release/config"
	"github.com/artyomtkachenko/release/meta"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateTmpDir(t *testing.T) {
	conf := config.Config{
		DataDir: "/tmp",
	}

	var testData meta.ReleaseMeta
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

	tmpPath, err := CreateTmpDir(m, conf)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	fullPath := filepath.Join(conf.DataDir, tmpPath.Path)
	f, _ := os.Open(filepath.Join(fullPath, "package.json"))

	if err := json.NewDecoder(f).Decode(&testData); err != nil {
		panic(err)
	}

	if testData.Project.Name != "foo" {
		t.Fatalf("Expected to get foo as a project name, but got: %v", testData.Project.Name)
	}

	defer os.Remove(fullPath)
}
