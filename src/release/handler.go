package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"release/extractor"
	"release/packer"
	"release/packer/rpm"
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
func DoBuild(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	packageType := vars["packageType"]

	if packageType != "rpm" {
		ThrowError(w, 400, "Only rpm is supported at the moment")
		return
	}
	req.ParseMultipartForm(0) //Do not use RAM, use tmp files to store streams
	releaseMeta, err := extractor.ExtractMeta(req)
	if err != nil {
		ThrowError(w, 400, err.Error())
		return
	}
	releaseMeta.Package.Type = packageType

	//Validates input
	if err := validate.Input(releaseMeta); err != nil {
		ThrowError(w, 422, err.Error())
		return
	}

	//Creates temorary build directory
	uniqBuildId, err := packer.CreateTmpDir(Config.DataDir)
	if err != nil {
		ThrowError(w, 400, err.Error())
		return
	}

	buildRoot := filepath.Join(Config.DataDir, uniqBuildId)

	if packageType == "rpm" {
		rpm.GenerateRpmBuildDirs(buildRoot)
		buildDir := filepath.Join(buildRoot, "BUILD")

		if err := extractor.Extract(buildDir, req); err != nil {
			ThrowError(w, 400, err.Error())
			return
		}
		scripts, err := extractor.ExtractScripts(req)
		if err != nil {
			ThrowError(w, 400, err.Error())
			return
		}
		if err := rpm.GenerateRpmSpec(releaseMeta, buildRoot, scripts); err != nil {
			ThrowError(w, 400, err.Error())
			return
		}
		if err := rpm.RunRpmBuild(releaseMeta, buildRoot); err != nil {
			ThrowError(w, 400, err.Error())
			return
		}
		//publish
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
