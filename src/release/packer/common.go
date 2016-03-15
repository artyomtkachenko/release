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

func CreateTmpDir(rm meta.ReleaseMeta, conf config.Config) (string, error) {
	uniqNumber := rand.Intn(9000000000)
	buildPath := rm.Project.Name + strconv.Itoa(uniqNumber)
	fullPath := filepath.Join(conf.DataDir, buildPath)
	fullBuildPath := filepath.Join(fullPath, "BUILD")

	if err := os.MkdirAll(fullBuildPath, 0755); err != nil {
		return "", errors.New(err.Error())
	}

	f, err := os.Create(filepath.Join(fullPath, "package.json"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(rm); err != nil {
		panic(err)
	}

	return rm.Project.Name + strconv.Itoa(uniqNumber), nil
}
