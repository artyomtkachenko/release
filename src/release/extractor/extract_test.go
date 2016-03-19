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

func TestExtractSucess(t *testing.T) {
	var ret error
	filename := "testdata/HelloWorld.zip"
	// fp, _ := os.Open(filename)
	// fileInfo, _ := fp.Stat()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ret = Extract("testdata/build", r)
	}))
	defer ts.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, _ := bodyWriter.CreateFormFile("data", filename)
	fh, _ := os.Open(filename)
	defer fh.Close()
	io.Copy(fileWriter, fh)

	contentType := bodyWriter.FormDataContentType()
	/* fmt.Printf("%+v\n", bodyBuf) */
	bodyWriter.Close()

	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)
	testFiles := []string{
		filepath.Join("testdata", "build", "java", "HelloWorld.jar"),
		filepath.Join("testdata", "build", "python", "HelloWorld.py"),
		filepath.Join("testdata", "build", "ruby", "HelloWorld.rb"),
	}

	for _, file := range testFiles {

		if _, err := os.Open(file); err != nil {
			t.Errorf("Expected file %s to exist\n", file)
		}
	}
	defer os.RemoveAll("testdata/build")
}
