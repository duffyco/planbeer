package main

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/go-kivik/couchdb/v4"
	cmp "github.com/google/go-cmp/cmp"
	guuid "github.com/google/uuid"
)


func DBRecipeExists( xmlRec XmlRecipe ) bool {
	row := getDB( "recipes2").Get( context.TODO(), xmlRec.Name )
	
	return ( row.ContentLength != 0 );
}

func DBGetRecipe( name string ) (RecipeDBO, error ) {
	row := getDB( "recipes2").Get( context.TODO(), name )

	if ( row.ContentLength == 0 ) {
		return RecipeDBO{}, errors.New( "Recipe not found: " + name )
	}

	var  curRecipe RecipeDBO
	err := row.ScanDoc(&curRecipe);
	
	if( err != nil ){
		return RecipeDBO{}, errors.New( "Recipe bad format: " + name )
	}

	return curRecipe, nil
}


func DBDeleteRecipe( rec Recipe ) bool {

	recdbo, err := DBGetRecipe( rec.XmlRecipe.Name )

	if( err != nil ) {
		return false;
	}

	_, err = getDB( "recipes2").Delete( context.TODO(), rec.XmlRecipe.Name, recdbo.Rev )
		
	if( err != nil ){
		return false;
	}

	return true;
}

func DBGetRecipes( databaseName string ) []RecipeDBO  {

	var recipeList []RecipeDBO;
			
	rows, err := getDB("recipes2").AllDocs( context.TODO(), map[string]interface{}{"include_docs": true} )

	if( err != nil ) {
		return recipeList
	}

	for rows.Next() {
		var sdbo RecipeDBO;
		if err := rows.ScanDoc(&sdbo); err != nil {
			panic(err)
		}

		recipeList = append( recipeList, sdbo )
	}

	return recipeList;
}

func DBImportRecipe( recipeName string, recipeID int, recipe XmlRecipe ) bool {
	
	row := getDB( "recipes2").Get( context.TODO(), recipeName )

	if ( row.ContentLength == 0 ) {
		recipeDBO := &RecipeDBO { 
			ID: guuid.New().String(),
			XmlRecipe: recipe, 
			RecipeID: recipeID,
			Synced: false,
		}

		rev, err := getDB("recipes2").Put( context.TODO(), recipeName, recipeDBO )
		
		if( err != nil ) {
			panic( err )
		}

		fmt.Printf("Found New Recipe: %s.  Importing(%s)\n", recipeName, rev)
		return true;
	} else {
		var  curRecipe RecipeDBO
		err := row.ScanDoc(&curRecipe);
		
		if( err != nil ){
			panic( err )
		}

		if( !cmp.Equal( curRecipe.XmlRecipe, recipe ) ) {
			curRecipe.XmlRecipe = recipe
			rev, err := getDB("recipes2").Put( context.TODO(), recipeName, curRecipe )

			if( err != nil ) {
				panic( err )
			}		

			fmt.Printf("Found New Recipe: %s.  Importing(%s)\n", recipeName, rev)
			return true;
		}

		return false;
	}
}

func DBSetRecipe( recipe Recipe ) error {
	
	// Get existing recipe for rev 
	row := getDB( "recipes2").Get( context.TODO(), recipe.XmlRecipe.Name )

	if ( row.ContentLength == 0 ) {
		return errors.New( "Recipe not found: " + recipe.XmlRecipe.Name );
	}

	var  curRecipe RecipeDBO
	err := row.ScanDoc(&curRecipe);
	
	if( err != nil ){
		return err
	}

	updatedDbo := convertRecipeToDBRecipe( recipe )
	updatedDbo.Rev = curRecipe.Rev

	rev, err := getDB("recipes2").Put( context.TODO(), recipe.XmlRecipe.Name, updatedDbo )
		
	if( err != nil ) {
		return err;
	}

	fmt.Printf("Updated Recipe: %s.  Importing(%s)\n", recipe.XmlRecipe.Name, rev)
	return nil;
}

