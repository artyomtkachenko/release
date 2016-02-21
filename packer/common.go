package packer

import (
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
	fullBuildPath := filepath.Join(conf.DataDir, buildPath)

	if err := os.MkdirAll(fullBuildPath, 0755); err != nil {
		return TmpDir{Path: ""}, errors.New(err.Error())
	}

	return TmpDir{Path: rm.Project.Name + strconv.Itoa(uniqNumber)}, nil
}