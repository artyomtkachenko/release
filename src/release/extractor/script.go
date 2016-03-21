package extractor

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var (
	ErrUnknownScript        = errors.New("Only shell scripts (.sh) are supported")
	ErrScriptAlreadyDefined = errors.New("Multiple of the same script type defined")
)

func ExtractScripts(req *http.Request) (map[string]string, error) {
	scriptsTypes := []string{"pre", "post", "preun", "postun"}
	var result = make(map[string]string)

	for _, script := range scriptsTypes {
		if _, ok := result[script]; ok {
			return result, ErrScriptAlreadyDefined // Not sure how to test it. Golang does not allow multiple entries as curl does, And I do not want to use curl in tests
		}
		stream, header, err := req.FormFile(script)
		if err == http.ErrMissingFile {
			continue
		}
		if err != nil {
			return result, err
		}
		fileExt := filepath.Ext(header.Filename)
		if fileExt != ".sh" {
			return result, ErrUnknownScript
		}
		body, err := ioutil.ReadAll(stream)
		if err != nil {
			return result, err
		}
		result[script] = string(body)
		stream.Close()
	}
	return result, nil
}
