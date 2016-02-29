package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := Routes{
		Route{
			Name:        "DoInit",
			Method:      "POST",
			Pattern:     "/release/v1/init/{packageType}",
			HandlerFunc: DoInit,
		},
		Route{
			Name:        "DoBuild",
			Method:      "PUT",
			Pattern:     "/release/v1/build/{buildId}",
			HandlerFunc: DoBuild,
		},
	}

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
