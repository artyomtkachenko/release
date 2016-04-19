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

	"github.com/artyomtkachenko/release/config"
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

func Test_getBuildDir(t *testing.T) {
	if res := getBuildDir("foo"); res != filepath.Join("foo", "BUILD") {
		t.Error("Failed returning the right path")
	}
}
func Test_specBuildDir(t *testing.T) {
	if res := getSpecDir("foo"); res != filepath.Join("foo", "SPEC") {
		t.Error("Failed returning the right path")
	}
}
func Test_rpmsBuildDir(t *testing.T) {
	if res := getRpmsDir("foo"); res != filepath.Join("foo", "RPMS") {
		t.Error("Failed returning the right path")
	}
}

func TestGenerateRpmSpec(t *testing.T) {
	m := config.ReleaseConfig{
		Project: config.Project{
			Name:        "foo",
			ContentRoot: "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Publish: config.Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: config.Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: config.Deploy{
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
	rm := config.ReleaseConfig{
		Project: config.Project{
			Name:        "foo",
			ContentRoot: "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Publish: config.Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: config.Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: config.Deploy{
			User:    "bob",
			Group:   "bob",
			RootDir: "/root",
		},
	}

	if err := GenerateRpmBuildDirs(testDataDir, rm); err != nil {
		t.Errorf("Could not generate RPM build dirs. Got error %+v\n", err)
	}
	buildDirs := []string{"BUILD", "SPEC", "RPMS"}
	for _, bd := range buildDirs {
		dir := filepath.Join(testDataDir, bd)
		_, err := os.Open(dir)
		if err != nil {
			t.Errorf("%+v\n", err)
		}
	}
	defer os.RemoveAll(testDataDir)
}

// func Test_convertBuildDirToDepoyDir(r *testing.T) {
// 	result := convertBuildDirToDepoyDir("Test_convertBuildDirToDepoyDir", "/opt/myapp", files)
// }
