package extractor

import (
	"errors"
	"net/http"
	"path/filepath"
)

var ErrUnknownFile = errors.New("unknown file type: valid file types are .zip, .tar")

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
		return ErrUnknownFile
	}
	defer stream.Close()
	return nil
}
