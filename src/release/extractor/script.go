package extractor

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func ExtractScripts(req *http.Request) (map[string]string, error) {
	scriptsTypes := []string{"post", "preun"}
	var result = make(map[string]string)

	for _, script := range scriptsTypes {
		stream, header, err := req.FormFile(script)
		if err == http.ErrMissingFile {
			continue
		}
		if err != nil {
			return result, err
		}
		fileExt := filepath.Ext(header.Filename)
		if fileExt != ".sh" {
			return result, errors.New("Only shell scripts are supported, got " + header.Filename)
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
