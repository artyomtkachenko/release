package packer

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"release/config"
	"release/meta"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateBuildDir(rm meta.ReleaseMeta, buildRoot string) (string, error) {
	if rm.Package.Type == "rpm" {
		return filepath.Join(buildRoot, "BUILD", rm.Deploy.RootDir), nil
	}
	return "", errors.New("Only rpm is supported at the moment") //TODO move error  to vars
}

func CreateTmpDir(rm meta.ReleaseMeta, conf config.Config) (string, error) {
	uniqNumber := rand.Intn(9000000000)
	buildName := strconv.Itoa(uniqNumber)
	buildRoot := filepath.Join(conf.DataDir, buildName)
	buildDir, err := GenerateBuildDir(rm, buildRoot)

	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return "", errors.New(err.Error())
	}

	f, err := os.Create(filepath.Join(buildRoot, "release.json"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(rm); err != nil {
		panic(err)
	}

	return buildName, nil
}
