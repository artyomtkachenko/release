package main

import (
	"log"
	"log/syslog"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logwriter, e := syslog.New(syslog.LOG_NOTICE, "release")
		if e == nil {
			log.SetOutput(logwriter)
		}

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
