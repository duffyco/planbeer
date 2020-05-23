package main

import (
	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Query	    string
	Value	    string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func init() {
	log.Printf("STARTING Routes: 12")
	for _, route := range routes {
		log.Printf("Route Added %s (%s) - %s", route.Name, route.Method, route.Pattern)
	}
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		"",
		"",
		Index,
	},
	Route{
		"ZState",
		"PUT",
		"/Vendors/input.cshtml",
		"type",
		"ZState",
		PicoBrewZState,
	},
	Route{
		"RecipeRefListController",
		"POST",
		"/Vendors/input.cshtml",
		"ctl",
		"RecipeRefListController",
		PicoBrewRecipeRefListController,
	},
	Route{
		"RecipeController",
		"Get",
		"/Vendors/input.cshtml",
		"type",
		"Recipe",
		PicoBrewRecipeController,
	},
	Route{
		"SessionController",
		"POST",
		"/Vendors/input.cshtml",
		"type",
		"ZSession",
		PicoBrewSessionController,
	},
	Route{
		"SessionLogController",
		"POST",
		"/Vendors/input.cshtml",
		"type",
		"ZSessionLog",
		PicoBrewSessionLogController,
	},
	/*
		Route{
			"ZSession",
			"POST",
			"/Vendors/input.cshtml?type=ZSession",
			PicoBrewZSession,
		},
		Route{
			"ZSessionLog",
			"POST",
			"/Vendors/input.cshtml?type=ZSessionLog",
			PicoBrewZSessionLog,
		},
		Route{
			"ZSessionLog",
			"POST",
			"/Vendors/input.cshtml?type=ZSessionLog",
			PicoBrewZSessionLog,
		},
	*/
}
