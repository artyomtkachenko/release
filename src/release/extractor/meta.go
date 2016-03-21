package extractor

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"release/config"
)

func ExtractMeta(req *http.Request) (config.ReleaseConfig, error) {
	var releaseMeta config.ReleaseConfig
	var metaData config.MetaData
	stream, _, err := req.FormFile("config")
	defer stream.Close()
	if err != nil {
		return releaseMeta, err
	}
	body, err := ioutil.ReadAll(stream)
	if err != nil {
		return releaseMeta, err
	}
	if err := json.Unmarshal(body, &releaseMeta); err != nil {
		return releaseMeta, err
	}
	meta := req.FormValue("meta")

	if err := json.Unmarshal([]byte(meta), &metaData); err != nil {
		return releaseMeta, err
	}
	releaseMeta.Project.Version = metaData.Version
	releaseMeta.Project.Revision = metaData.Revision

	return releaseMeta, nil
}
