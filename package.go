package main

import (
	"errors"
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

func CreateTmpDir(rm meta.ReleaseMeta) (TmpDir, error) {
	uniqNumber := rand.Intn(9000000000)
	buildPath := rm.Name + strconv.Itoa(uniqNumber)
	fullBuildPath := filepath.Join(Conf.DataDir, buildPath)

	if err := os.MkdirAll(fullBuildPath, 0755); err != nil {
		return TmpDir{Path: ""}, errors.New(err.Error())
	}

	return TmpDir{Path: rm.Name + strconv.Itoa(uniqNumber)}, nil
}
