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

func TestExtractScriptsSucess(t *testing.T) {
	var ret error
	filename := "testdata/simple.sh"
	testData := "TestExtractScriptsSucess"
	os.MkdirAll(filepath.Join(testData, "SCRIPTS"), 0755)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ret = ExtractScripts("TestExtractScriptsSucess", r)
	}))
	defer ts.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, _ := bodyWriter.CreateFormFile("post", filename)
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
		filepath.Join(testData, "SCRIPTS", "post.sh"),
	}

	for _, file := range testFiles {
		if _, err := os.Open(file); err != nil {
			t.Errorf("Expected file %s to exist\n", file)
		}
	}
	defer os.RemoveAll(testData)
}
