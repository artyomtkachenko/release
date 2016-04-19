package extractor

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

func ExtractZipPackage(buildDir string, stream multipart.File, size int64) error {
	body, err := ioutil.ReadAll(io.LimitReader(stream, size))
	if err != nil {
		return err
	}

	archive, err := zip.NewReader(bytes.NewReader(body), size)
	if err != nil {
		return err
	}

	for _, zf := range archive.File {
		path := filepath.Join(buildDir, zf.Name)

		if zf.FileInfo().IsDir() {
			os.MkdirAll(path, zf.Mode())
			continue
		}

		fileReader, err := zf.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}
