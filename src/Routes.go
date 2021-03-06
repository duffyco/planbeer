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
	log.Printf("STARTING Planbeer Server: " + VERSION)
	for _, route := range routes {
		log.Printf("Route Added %s (%s) - %s", route.Name, route.Method, route.Pattern)
	}
}

var routes = Routes{
	Route{
		"Index",
		"Get",
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
	Route{
		"UI-RecipeList",
		"GET",
		"/ui/ListRecipes",
		"",
		"",
		ListRecipes,
	},
	Route{
		"UI-MachineList",
		"GET",
		"/ui/ListMachines",
		"",
		"",
		ListMachines,
	},
	Route{
		"UI-SessionList",
		"GET",
		"/ui/ListSessions",
		"",
		"",
		ListSessions,
	},
	Route{
		"UI-GetSession",
		"GET",
		"/ui/GetSession/{id}",
		"",
		"",
		GetSessionHandler,
	},
	Route{
		"UI-DeleteSession",
		"GET",
		"/ui/DeleteSession/{id}",
		"",
		"",
		DeleteSessionHandler,
	},
	Route{
		"UI-GetRecipe",
		"GET",
		"/ui/GetRecipe/{id}",
		"",
		"",
		GetRecipe,
	},
	Route{
		"UI-DeleteRecipe",
		"GET",
		"/ui/DeleteRecipe/{id}",
		"",
		"",
		DeleteRecipeHandler,
	},
	Route{
		"UI-GetLogs",
		"GET",
		"/ui/GetLogs/{id}",
		"",
		"",
		GetLogs,
	},
	Route{
		"UI-GetImportableItems",
		"GET",
		"/ui/GetImportableItems",
		"",
		"",
		GetImportableItems,
	},
	Route{
		"UI-GetLogs",
		"GET",
		"/ui/GetLogs/{id}",
		"",
		"",
		GetLogs,
	},
	Route{
		"UI-ImportSession",
		"POST",
		"/ui/ImportSession/",
		"",
		"",
		ImportSession,
	},	
	Route{
		"UI-ImportRecipe",
		"POST",
		"/ui/ImportRecipe/{name}",
		"",
		"",
		ImportRecipe,
	},
	Route{
		"UI-DeleteImportableSession",
		"GET",
		"/ui/DeleteImportableSession/{id}",
		"",
		"",
		DeleteImportableSessionHandler,
	},	
	Route{
		"UI-DeleteImportableRecipe",
		"GET",
		"/ui/DeleteImportableRecipe/{id}",
		"",
		"",
		DeleteImportableRecipeHandler,
	},
	Route{
		"UI-ImportSession",
		"POST",
		"/ui/UploadFiles/",
		"",
		"",
		UploadFiles,
	},	
	Route{
		"UI-UpdateRecipe",
		"POST",
		"/ui/UpdateRecipe/{id}",
		"",
		"",
		UpdateRecipeFromHeader,
	},			
	Route{
		"UI-DeleteMachine",
		"GET",
		"/ui/DeleteMachine/{token}",
		"",
		"",
		DeleteMachineHandler,
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
