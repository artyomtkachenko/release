package extractor

import (
	"net/http"
	"net/http/httptest"
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
	contentType, bodyBuf := SendMultiPart(filename, "post")

	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)

	if !strings.Contains(scripts["post"], "#!/bin/sh") {
		t.Errorf("Expected to get #!/bin/sh, got %+v", scripts)
	}
}

func TestExtractFailErrErrUnknownScript(t *testing.T) {
	var ret error
	filename := "testdata/release.json"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ret = ExtractScripts(r)
	}))
	defer ts.Close()

	contentType, bodyBuf := SendMultiPart(filename, "post")
	req, _ := http.NewRequest("POST", ts.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	client.Do(req)
	if ret != ErrUnknownScript {
		t.Error("It should fail")
	}
}
