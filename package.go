package main

type tmpDir struct {
	Id string `json:"id"`
}

//creates temporarty directory for the future build
// Should create the following structure:
//
// /data/release/uniqhash/meta.json the metadatto build an rpm
//
// /data/release/uniqhash/build to store the built rpm
//
// /data/release/uniqhash/package to store the content of the ZIP package provided by a client side application e.g. curl
func CreateTmpDir(releaseMeta ReleaseMeta) tmpDir {
	return tmpDir
}
