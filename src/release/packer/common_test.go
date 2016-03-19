package packer

import (
	"os"
	"testing"
)

func TestCreateTmpDir(t *testing.T) {
	testData := "TestCreateTmpDir"
	_, err := CreateTmpDir(testData)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	defer os.Remove(testData)
}
