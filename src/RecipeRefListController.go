package main

import (
	"log"
	"strconv"
)

func createRecipeRefListControllerRespMsg( inmsg RecipeRefListControllerMsg ) RecipeRefListControllerRespMsg {
	
	var retValue RecipeRefListControllerRespMsg 
	retValue.Kind = inmsg.Kind;
	retValue.Offset = inmsg.Offset;
	retValue.SearchString = nil;
	retValue.TotalResults = len( recipeList )

	if( retValue.TotalResults > inmsg.MaxCount ) {
		retValue.TotalResults = inmsg.MaxCount
	}
	
	for i, recipe := range recipeList {

		retValue.Recipes = append( retValue.Recipes, GetControllerRecipeRefDescription( recipe ) )

		log.Printf("Recipe Added %s (%d)", retValue.Recipes[i].Name, retValue.Recipes[i].ID)
	} 

	// @TODO: Setting this to 0 for some reason???
	retValue.MaxCount = 0

		//@TODO: Need to have a synced flag in metadata
		//@TODO: Kind variable needs to be adjusted based on type (Beer/Coffee/etc)

	return retValue
}

func GetControllerRecipeRefDescription(rec Recipe) RecipeDescription {
	var ret RecipeDescription

	ret.Abv = int(rec.XmlRecipe.ABV)
	ret.ID = rec.RecipeID
	ret.Ibu = rec.XmlRecipe.IBU
	ret.Kind, _ = strconv.Atoi(rec.XmlRecipe.Type);
	ret.Name = rec.ID
	ret.Uri = nil

	return ret
}