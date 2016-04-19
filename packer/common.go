package packer

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateTmpDir(dataDir string) (string, error) {
	uniqNumber := rand.Intn(9000000000)
	buildName := strconv.Itoa(uniqNumber)
	buildRoot := filepath.Join(dataDir, buildName)

	if err := os.MkdirAll(buildRoot, 0755); err != nil {
		return "", err
	}
	return buildName, nil
}
