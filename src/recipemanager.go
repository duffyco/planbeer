package main

import (
	"os"
	"io/ioutil"
	"log"
)

var recipeList []Recipe
var xmlRecipes []XmlRecipe  

func NewRecipeManager( importPath string ) {
	xmlFiles := ScanForRecipes( importPath )

	//Load Recipes
    for i, xmls := range xmlFiles {
		log.Printf( "Reading file: %s/%s", importPath, xmls.Name() )
		parsedStruct := RecipeFile( importPath + "/" + xmls.Name() )
		xmlRecipes = append( xmlRecipes, parsedStruct )
		recipeList = append( recipeList, ConvertToRecipe( i, parsedStruct ) )
	}
}

func ScanForRecipes( importPath string ) []os.FileInfo {
	files, err := ioutil.ReadDir(importPath)
	if err != nil {
        panic(err)
	}
	
	return files
}