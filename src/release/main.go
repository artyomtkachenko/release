package main

import (
	"flag"
	"log"
	"net/http"
	"release/config"
)

var Config config.Config

func main() {
	port := flag.String("port", ":8080", "A default port")
	dataDir := flag.String("data-dir", "/data/release", "Root dir for a server ")
	logFile := flag.String("log-file", "/var/log/release.log", "A default log file")
	flag.Parse()

	Config = config.Config{
		Port:    *port,
		LogFile: *logFile,
		DataDir: *dataDir,
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(Config.Port, router))
}
