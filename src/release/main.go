package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"release/config"
)

//

var (
	Config  config.ServerConfig
	version string
)

func main() {
	port := flag.String("port", ":8080", "A default port")
	dataDir := flag.String("data-dir", "/data/release", "Root dir for a server ")
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			fmt.Println(version)
			os.Exit(0)
		}
	}
	flag.Parse()

	Config = config.ServerConfig{
		Port:    *port,
		DataDir: *dataDir,
	}

	router := NewRouter()
	fmt.Fprintf(os.Stderr, "Please use localhost:%s\n", *port)
	log.Fatal(http.ListenAndServe(Config.Port, router))
}
