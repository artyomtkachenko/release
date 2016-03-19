package extractor

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func ExtractScripts(buildDir string, req *http.Request) error {
	scriptsTypes := []string{"post", "preun"}

	for _, script := range scriptsTypes {
		stream, header, err := req.FormFile(script)
		if err != nil {
			return err
		}
		fileExt := filepath.Ext(header.Filename)
		if fileExt != ".sh" {
			return errors.New("Only shell scripts are supported, got " + header.Filename)
		}
		body, err := ioutil.ReadAll(stream)
		if err != nil {
			return err
		}

		targetFile, err := os.Create(filepath.Join(buildDir, "SCRIPTS", script+".sh"))
		if _, err := targetFile.Write(body); err != nil {
			return err
		}
		targetFile.Close()

		stream.Close()
	}
	return nil
}
