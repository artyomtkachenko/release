package extractor

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func SendMultiPart(filename, field string) (string, *bytes.Buffer) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, _ := bodyWriter.CreateFormFile(field, filename)
	fh, _ := os.Open(filename)
	defer fh.Close()
	io.Copy(fileWriter, fh)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	return contentType, bodyBuf
}

func TestExtractSucess(t *testing.T) {
	filename := "testdata/HelloWorld.zip"

	testDir := "TestExtractSucess"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = Extract(testDir, r)
	}))
	defer ts.Close()

	contentType, bodyBuf := SendMultiPart(filename, "data")
	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)
	testFiles := []string{
		filepath.Join(testDir, "java", "HelloWorld.jar"),
		filepath.Join(testDir, "python", "HelloWorld.py"),
		filepath.Join(testDir, "ruby", "HelloWorld.rb"),
	}

	for _, file := range testFiles {

		if _, err := os.Open(file); err != nil {
			t.Errorf("Expected file %s to exist\n", file)
		}
	}
	defer os.RemoveAll(testDir)
}

func TestExtractFailErrUnknownFile(t *testing.T) {
	var ret error
	filename := "testdata/simple.sh"

	testDir := "TestExtractSucess"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ret = Extract(testDir, r)
	}))
	defer ts.Close()

	contentType, bodyBuf := SendMultiPart(filename, "data")
	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)
	if ret != ErrUnknownFile {
		t.Error("It should fail")
	}
}

func TestExtractFailFormFile(t *testing.T) {
	var ret error
	filename := "testdata/simple.sh"

	testDir := "TestExtractSucess"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ret = Extract(testDir, r)
	}))
	defer ts.Close()

	contentType, bodyBuf := SendMultiPart(filename, "nodata")
	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)
	if ret == nil {
		t.Error("It should fail")
	}
}
