package webdav

//publish to webdav endpoint

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func putFile(filename, targetUrl, username, password string) (string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	check(err)

	// open file handle
	fh, err := os.Open(filename)
	check(err)

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	check(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest("PUT", targetUrl, bodyBuf)
	check(err)
	req.Header.Set("Content-Type", contentType)
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)

	return resp.Status, nil
}
