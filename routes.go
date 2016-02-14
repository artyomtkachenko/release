package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"DoInit",
		"POST",
		"/v1/init/rhel",
		DoInit,
	},
	Route{
		"DoBuild",
		"POST",
		"/v1/build/{buildId}",
		DoBuild,
	},
}
