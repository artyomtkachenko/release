package extractor

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestExtractScriptsSucess(t *testing.T) {
	var ret error
	var scripts map[string]string
	filename := "testdata/simple.sh"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scripts, ret = ExtractScripts(r)
	}))
	defer ts.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, _ := bodyWriter.CreateFormFile("post", filename)
	fh, _ := os.Open(filename)
	defer fh.Close()
	io.Copy(fileWriter, fh)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)

	if !strings.Contains(scripts["post"], "#!/bin/sh") {
		t.Errorf("Expected to get #!/bin/sh, got %+v", scripts)
	}
}
