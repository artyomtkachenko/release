package main
import (
	"os"
	"math/rand"
	"time"
	"strconv"
	"errors"
)

type tmpDir struct {
	Path string `json:"path"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateTmpDir(rm ReleaseMeta) (tmpDir, error) {
  uniqNumber := rand.Intn(9000000000)
	buildPath := "C:\\release\\" + rm.Name + strconv.Itoa(uniqNumber)

  if err := os.Mkdir(buildPath, 0755); err != nil {
    return tmpDir{Path: ""}, errors.New(err.Error())
	}

	return tmpDir{Path: buildPath}, nil
}
