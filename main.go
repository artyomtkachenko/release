package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	Conf Config
)

func main() {
	port := flag.String("port", ":8080", "A default port")
	dataDir := flag.String("data-dir", "/data/release", "Root dir for a server ")
	logFile := flag.String("log-file", "/var/log/release.log", "A default log file")
	flag.Parse()

	Conf = Config{
		Port:    *port,
		LogFile: *logFile,
		DataDir: *dataDir,
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(Conf.Port, router))
}
