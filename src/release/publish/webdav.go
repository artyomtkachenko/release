package webdav

//publish to webdav endpoint

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func postFile(filename, targetUrl, username, password string) ([]byte, error) {
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
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)

	resp_body, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return resp_body, nil
}
