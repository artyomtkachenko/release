package packer

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"release/config"
	"release/meta"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateTestData(path string) {
	testData := []string{"bin", "etc", "log", "data"}
	for _, f := range testData {
		fp := filepath.Join(path, "BUILD", f)
		os.MkdirAll(fp, 0755)
		for i := 0; i <= rand.Intn(5); i++ {
			f, _ := os.Create(filepath.Join(fp, "test_"+strconv.Itoa(i)+".txt"))
			f.Close()
		}
	}
}

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
	generateTestData("/tmp/aaa")
	err := GenerateRpmSpec(m, conf, tmpId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	defer os.RemoveAll("/tmp/aaa")
}
