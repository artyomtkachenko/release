package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"release/extractor"
	"release/meta"
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
	var buildDir string //Package specific dir

	if packageType != "rpm" {
		ThrowError(w, 400, "Only rpm is supported at the moment")
		return
	}
	req.ParseMultipartForm(0) //Do not use RAM, use tmp files to store streams
	releaseMeta, err := meta.ExtractMeta(req)
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
		if err := extractor.ExtractScripts(buildDir, req); err != nil {
			ThrowError(w, 400, err.Error())
			return
		}
		//generateSpec
		//run rpm build
		//publish
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

//Gets ZIP stream and extracts it under the predefined location
func DoBuild(w http.ResponseWriter, req *http.Request) {
	var releaseMeta meta.ReleaseMeta
	vars := mux.Vars(req)
	buildId := vars["buildId"]
	packageType := vars["packageType"]
	buildRoot := filepath.Join(Config.DataDir, buildId)

	//Reading the release.json config file
	releaseConfig := filepath.Join(buildRoot, "release.json")
	releaseConfigBody, err := ioutil.ReadFile(releaseConfig)
	if err != nil {
		ThrowError(w, 400, err.Error())
		return
	}

	if err := json.Unmarshal(releaseConfigBody, &releaseMeta); err != nil {
		ThrowError(w, 422, err.Error())
		return
	}

	buildDir, _ := packer.GenerateBuildDir(releaseMeta, buildRoot)

	//Verify we have a build directory first
	if _, err := os.Stat(buildDir); err != nil {
		ThrowError(w, 400, err.Error())
		return
	}

	if releaseMeta.Package.Type == "rpm" && packageType == "rpm" {
		// Now we can finally generate the RPM Spec file
		if err := packer.GenerateRpmSpec(releaseMeta, buildRoot); err != nil {
			ThrowError(w, 400, err.Error())
			return
		}
		if err := packer.RunRpmBuild(releaseMeta, buildRoot); err != nil {
			ThrowError(w, 400, err.Error())
			return
		}
	}
	/* defer os.RemoveAll(buildRoot) */
}
