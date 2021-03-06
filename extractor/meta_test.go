package extractor

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/artyomtkachenko/release/config"
)

func TestExtractMetaSuccess(t *testing.T) {
	var rm config.ReleaseConfig
	meta := `{"version": "1.0.0", "revision": "abd1767a214"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rm, _ = ExtractMeta(r)
	}))
	defer ts.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("meta", meta)

	filename := "testdata/release.json"
	fileWriter, _ := bodyWriter.CreateFormFile("config", filename)
	fh, _ := os.Open(filename)
	defer fh.Close()
	defer bodyWriter.Close()
	io.Copy(fileWriter, fh)

	contentType := bodyWriter.FormDataContentType()
	/* fmt.Printf("%+v\n", bodyBuf) */
	bodyWriter.Close()

	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)

	if rm.Project.Name != "activemq" ||
		rm.Project.Version != "1.0.0" ||
		rm.Project.Revision != "abd1767a214" {
		t.Errorf("Expected Name: activemq, Version: 1.0.0, Revision: abd1767a214\nGot:%+v\n", rm)
	}
}
