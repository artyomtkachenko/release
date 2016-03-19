package extractor

import (
	"errors"
	"net/http"
	"path/filepath"
)

func Extract(buildDir string, req *http.Request) error {
	stream, header, err := req.FormFile("data")
	if err != nil {
		return err
	}

	fileExt := filepath.Ext(header.Filename)
	if fileExt == ".zip" {
		if ExtractZipPackage(buildDir, stream, req.ContentLength); err != nil {
			return err
		}
	} else {
		return errors.New("Recieved unknown file type: valid type is .zip, got " + fileExt)
	}
	defer stream.Close()
	return nil
}
