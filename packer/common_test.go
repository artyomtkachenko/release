package packer

import (
	"path/filepath"
	"testing"
)

//Config mock
type Config struct {
	DataDir string
}

func TestCreateTmpDir(t *testing.T) {
	Conf := Config{
		DataDir: "/tmp",
	}

	meta := meta.ReleaseMeta{
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

	tmpPath, err := CreateTmpDir(meta)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	fullPath := filepath.Join(Config.DataDir, tmpPath.Path)
	defer os.Remove(fullPath)
}
