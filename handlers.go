package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
)

func ThrowError(w http.ResponseWriter, code int, text string)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	e := jsonErr{
		Code: code,
		Text: text,
	}
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}
//Initialize the build.
func DoInit(w http.ResponseWriter, req *http.Request) {
	var releaseMeta ReleaseMeta
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, req.ContentLength))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &releaseMeta); err != nil {
	  ThrowError(w, 422, err.Error())
	}
 //TODO KILLME
	fmt.Printf("%+v\n", releaseMeta)
  //Validates input
	if err := validateInput(releaseMeta); err != nil {
    ThrowError(w, 422, err.Error())
		return
	}

  //Creates temorary build directory
	path, err := CreateTmpDir(releaseMeta)
	if err != nil {
		ThrowError(w, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(path); err != nil {
		panic(err)
	}
}

//Gets ZIP file and extracts it under the predefined location
func DoBuild(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	r, err := zip.NewReader(bytes.NewReader(body), req.ContentLength)
	if err != nil {
		panic(err)
	}

	for _, zf := range r.File {
		dst, err := os.Create(zf.Name)
		if err != nil {
			panic(err)
		}
		defer dst.Close()
		src, err := zf.Open()
		if err != nil {
			panic(err)
		}
		defer src.Close()

		io.Copy(dst, src)
	}
}
