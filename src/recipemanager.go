package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var recipeList []Recipe

func getRecipes() []Recipe {
	return recipeList;
}

func getRecipe(i int) Recipe {
	return recipeList[i];
}

func setRecipe( i int, recipe Recipe ) ( error ) {
	if( recipeList[i].ID != recipe.ID ) {
		return errors.New( "Error saving recipe: " + recipeList[i].ID + " != " + recipe.ID)
	}

	err := DBSetRecipe( recipe )

	if( err != nil ) {
		return err;
	}

	loadRecipes();
	return nil;
}

func loadRecipes() {
	dbos:= DBGetRecipes( "recipes2" );

	var maxID int = -1;
	for _,recipe := range dbos {
		if( recipe.RecipeID > maxID ) {
			maxID = recipe.RecipeID
		}
	}

	recipeList = make( []Recipe, maxID + 1 )

	for _, recipe := range dbos {
		recipeList[recipe.RecipeID] = convertDBRecipeToRecipe( recipe )
	}

	log.Printf( "Loaded %d Recipes", maxID + 1 );
}

func NewRecipeManager( importPath string ) {
	loadRecipes();
	xmlFiles := ScanForRecipes( importPath )

	//Load Recipes
    for _, xmls := range xmlFiles {
		log.Printf( "Reading file: %s/%s", importPath, xmls.Name() )
		parsedStruct, err := RecipeFile( importPath + "/" + xmls.Name() )
		if( err != nil ) {
			continue;
		}
		
		recID := len( recipeList );
		if( DBImportRecipe( parsedStruct.Name, recID, parsedStruct ) ) {
			recipeList = append( recipeList, ConvertToRecipe( recID, parsedStruct ) )
		}
	}
}

func DeleteRecipe( id int ) bool {
	if( id > len( recipeList ) ) {
		return false
	}
	
	ret := DBDeleteRecipe( getRecipe( id ) );
	
	loadRecipes();
	return ret;
}

func UpdateRecipe( newRecipe XmlRecipe ) bool{
	recID := len( recipeList );
	if( !DBImportRecipe( newRecipe.Name, recID, newRecipe ) ) {
		return false;
	}

	loadRecipes();
	return true;
}

func ScanForRecipes( importPath string ) []os.FileInfo {
	files, err := ioutil.ReadDir(importPath)
	if err != nil {
        panic(err)
	}
	
	return files
}

