package packer

import (
	"encoding/json"
	"errors"
	"github.com/artyomtkachenko/release/config"
	"github.com/artyomtkachenko/release/meta"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type TmpDir struct {
	Path string `json:"path"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateTmpDir(rm meta.ReleaseMeta, conf config.Config) (TmpDir, error) {
	uniqNumber := rand.Intn(9000000000)
	buildPath := rm.Project.Name + strconv.Itoa(uniqNumber)
	fullPath := filepath.Join(conf.DataDir, buildPath)
	fullBuildPath := filepath.Join(fullPath, "BUILD")

	if err := os.MkdirAll(fullBuildPath, 0755); err != nil {
		return TmpDir{Path: ""}, errors.New(err.Error())
	}

	f, err := os.Create(filepath.Join(fullPath, "package.json"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(rm); err != nil {
		panic(err)
	}

	return TmpDir{Path: rm.Project.Name + strconv.Itoa(uniqNumber)}, nil
}
