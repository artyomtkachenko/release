package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/artyomtkachenko/release/meta"
	"github.com/artyomtkachenko/release/packer"
	"github.com/artyomtkachenko/release/validate"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func ThrowError(w http.ResponseWriter, code int, text string) {
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
	var releaseMeta meta.ReleaseMeta
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
	fmt.Printf("%+v\n", Config)
	fmt.Printf("%+v\n", releaseMeta)
	//Validates input
	if err := validate.Input(releaseMeta); err != nil {
		ThrowError(w, 422, err.Error())
		return
	}

	//Creates temorary build directory
	path, err := packer.CreateTmpDir(releaseMeta, Config)
	if err != nil {
		ThrowError(w, 500, err.Error())
		return
	}

	if releaseMeta.Package.Type == "rpm" {
		err := packer.ConvertJSON2RpmSpec(releaseMeta, Config, path)
		if err != nil {
			panic(err)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(path); err != nil {
		panic(err)
	}
}

//Gets ZIP stream and extracts it under the predefined location
func DoBuild(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	buildId := vars["buildId"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	archive, err := zip.NewReader(bytes.NewReader(body), req.ContentLength)
	if err != nil {
		panic(err)
	}

	for _, zf := range archive.File {
		dst := filepath.Join(Config.DataDir, buildId, "BUILD")
		path := filepath.Join(dst, zf.Name, ".spec")

		if zf.FileInfo().IsDir() {
			os.MkdirAll(path, zf.Mode())
			continue
		}

		fileReader, err := zf.Open()
		if err != nil {
			panic(err)
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if err != nil {
			panic(err)
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			panic(err)
		}
	}
}
