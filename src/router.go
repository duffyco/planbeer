package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Queries( route.Query, route.Value).
			Handler(handler)

	}
	return router
}
