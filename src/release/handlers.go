package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"release/meta"
	"release/packer"
	"release/validate"

	"github.com/gorilla/mux"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

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
	vars := mux.Vars(req)
	packageType := vars["packageType"]

	if packageType != "rpm" {
		ThrowError(w, 400, "Only rpm is supported at the moment")
	}

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, req.ContentLength))
	check(err)
	err = req.Body.Close()
	check(err)

	if err := json.Unmarshal(body, &releaseMeta); err != nil {
		ThrowError(w, 422, err.Error())
	}
	//Validates input
	if err := validate.Input(releaseMeta); err != nil {
		ThrowError(w, 422, err.Error())
		return
	}

	//Creates temorary build directory
	uniqBuildId, err := packer.CreateTmpDir(releaseMeta, Config)
	if err != nil {
		ThrowError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uniqBuildId)
	check(err)
}

//Gets ZIP stream and extracts it under the predefined location
func DoBuild(w http.ResponseWriter, req *http.Request) {
	var releaseMeta meta.ReleaseMeta
	vars := mux.Vars(req)
	buildId := vars["buildId"]
	buildRoot := filepath.Join(Config.DataDir, buildId)
	buildDir := filepath.Join(buildRoot, "BUILD")

	//Verify we have a build directory first
	if _, err := os.Stat(buildDir); err != nil {
		ThrowError(w, 400, err.Error())
	}

	//Reading the whole ZIP package into the RAM. TODO benchmark it and check if it is possible to read it in chunks
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ThrowError(w, 400, err.Error())
	}

	archive, err := zip.NewReader(bytes.NewReader(body), req.ContentLength)
	if err != nil {
		ThrowError(w, 400, err.Error())
	}

	for _, zf := range archive.File {

		path := filepath.Join(buildDir, zf.Name)

		if zf.FileInfo().IsDir() {
			os.MkdirAll(path, zf.Mode())
			continue
		}

		fileReader, err := zf.Open()
		if err != nil {
			ThrowError(w, 400, err.Error())
			panic(err)
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if err != nil {
			ThrowError(w, 400, err.Error())
			panic(err)
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			ThrowError(w, 400, err.Error())
			panic(err)
		}
	}

	//Reading the releas.json config file
	releaseConfig := filepath.Join(buildRoot, "release.json")
	releaseConfigBody, err := ioutil.ReadFile(releaseConfig)
	if err != nil {
		ThrowError(w, 400, err.Error())
	}

	if err := json.Unmarshal(releaseConfigBody, &releaseMeta); err != nil {
		ThrowError(w, 422, err.Error())
	}

	if releaseMeta.Package.Type == "rpm" {
		// Now we can finally generate the RPM Spec file
		if err := packer.GenerateRpmSpec(releaseMeta, buildRoot); err != nil {
			ThrowError(w, 400, err.Error())
		}
	}
}
