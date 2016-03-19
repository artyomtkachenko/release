package extractor

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"release/meta"
)

func ExtractMeta(req *http.Request) (meta.ReleaseMeta, error) {
	var releaseMeta meta.ReleaseMeta
	var metaData meta.MetaData
	config, _, err := req.FormFile("config")
	defer config.Close()
	if err != nil {
		return releaseMeta, err
	}
	body, err := ioutil.ReadAll(config)
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
