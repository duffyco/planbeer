package main

import (
	"log"
)

type RecipeRefListControllerMsg struct {
	Kind int `json:"Kind"`
	MaxCount      int    `json:"MaxCount"`
	Offset      int    `json:"Offset"`
}

type RecipeDescription struct {
	Abv int `json:"Abv"`
	ID int `json:"ID"`
	Ibu int `json:"Ibu"`
	Kind int `json:"Kind"`
	Name string `json:"Name"`
	Uri *string `json:"Uri"`
}

type RecipeRefListControllerRespMsg struct {
	Kind int `json:"Kind"`
	Offset      int    `json:"Offset"`
	SearchString *string `json:"SearchString"`
	MaxCount      int    `json:"MaxCount"`
	TotalResults      int    `json:"TotalResults"`
	Recipes 	[]RecipeDescription  `json:"Recipes"`
}

func init() {
	NewRecipeManager( "/recipes")
}

func createRecipeRefListControllerRespMsg( inmsg RecipeRefListControllerMsg ) RecipeRefListControllerRespMsg {
	
	var retValue RecipeRefListControllerRespMsg 
	retValue.Kind = inmsg.Kind;
	retValue.Offset = inmsg.Offset;
	retValue.SearchString = nil;
	retValue.TotalResults = len( recipeList )
	
	for i, recipe := range recipeList {

		retValue.Recipes = append( retValue.Recipes, GetRecipeDescription( recipe ) )

		log.Printf("Recipe Added %s (%d)", retValue.Recipes[i].Name, retValue.Recipes[i].ID)
	} 

//	if( length(retValue.Recipes) > inmsg.MaxCount )
//		retValue.MaxCount = inmsg.MaxCount
	// @TODO: Setting this to 0 for some reason???
	retValue.MaxCount = 0

	return retValue
}
