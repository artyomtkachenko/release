package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f, err := os.OpenFile(Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%d\t%s",
			r.Method,
			r.RequestURI,
			name,
			r.ContentLength,
			time.Since(start),
		)
	})
}
