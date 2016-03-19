package rpm

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"release/meta"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateTestData(path string) {
	testData := []string{"bin", "etc", "log", "data"}

	os.MkdirAll(filepath.Join(path, "SPEC"), 0755)
	for _, f := range testData {
		fp := filepath.Join(path, "BUILD", f)
		os.MkdirAll(fp, 0755)

		for i := 0; i <= rand.Intn(5); i++ {
			f, _ := os.Create(filepath.Join(fp, "test_"+strconv.Itoa(i)+".txt"))
			f.Close()
		}
	}
}

func TestGenerateRpmSpec(t *testing.T) {
	m := meta.ReleaseMeta{
		Project: meta.Project{
			Name:        "foo",
			ContentRoot: "/root",
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

	testDataDir := filepath.Join("TestGenerateRpmSpec")
	generateTestData(testDataDir)
	scripts := map[string]string{
		"post":  "#!/bin/sh\necho postInstall",
		"preun": "#!/bin/sh\necho 40054dea1ba2f1f425ef792e6cee80b4",
	}
	err := GenerateRpmSpec(m, testDataDir, scripts)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	spec, _ := ioutil.ReadFile(filepath.Join(testDataDir, "SPEC", "foo.spec"))
	if !strings.Contains(string(spec), "40054dea1ba2f1f425ef792e6cee80b4") {
		t.Errorf("Did not found the key string 40054dea1ba2f1f425ef792e6cee80b4 in the spec file")
	}

	defer os.RemoveAll(testDataDir)
}

func TestGenerateRpmBuildDirs(t *testing.T) {
	testDataDir := "TestGenerateRpmBuildDirs"
	if err := GenerateRpmBuildDirs(testDataDir); err != nil {
		t.Errorf("Could not generate RPM build dirs. Got error %+v\n", err)
	}
	defer os.RemoveAll(testDataDir)
}

// func Test_convertBuildDirToDepoyDir(r *testing.T) {
// 	result := convertBuildDirToDepoyDir("Test_convertBuildDirToDepoyDir", "/opt/myapp", files)
// }
