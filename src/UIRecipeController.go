package main

import "strconv"

func ListRecipesController() UIRecipeList {

	var uiRecipeList UIRecipeList
	loadRecipes();

	recList := getRecipes()

	for _, recipe := range recList {
		uiRecipeList.Recipes = append(uiRecipeList.Recipes, convertRecipetoRecipeList(recipe))
	}

	return uiRecipeList
}

func DeleteUIRecipe(id int) DefaultRespMsg {
	var respMsg DefaultRespMsg

	respMsg.Value = strconv.FormatBool(DeleteRecipe(id))
	return respMsg
}

func UpdateUIRecipe( id int, inmsg UpdateRecipeList ) DefaultRespMsg {
	var respMsg DefaultRespMsg

	updateRec := getRecipe( id )
	updateRec.Synced = inmsg.Synced

	err := setRecipe( id, updateRec )

	respMsg.Value = strconv.FormatBool( err == nil )
	return respMsg;
}

func ViewRecipe(id int) UIRecipeView {
	var retView UIRecipeView
	retView.Recipe = getRecipe(id).XmlRecipe
	return retView
}

type UIRecipeView struct {
	Recipe XmlRecipe `json:"Recipe"`
}

type UIRecipeList struct {
	Recipes []RecipeList `json:"Recipes"`
}

func convertRecipetoRecipeList(recipe Recipe) RecipeList {
	var retList RecipeList

	retList.ID = recipe.RecipeID
	retList.Name = recipe.XmlRecipe.Name
	retList.Style = recipe.XmlRecipe.Style.Name
	retList.ABV = recipe.XmlRecipe.ABV
	retList.FG = recipe.XmlRecipe.FG
	retList.OG = recipe.XmlRecipe.OG
	retList.IBU = recipe.XmlRecipe.IBU
	retList.SRM = recipe.XmlRecipe.Color
	retList.Synced = recipe.Synced

	return retList
}

type RecipeList struct {
	ID     int     `json:"ID"`
	Name   string  `json:"name"`
	Style  string  `json:"style"`
	OG     float32 `json:"og"`
	FG     float32 `json:"fg"`
	IBU    int     `json:"ibu"`
	SRM    int     `json:"srm"`
	ABV    float32 `json:"abv"`
	Synced bool    `json:"synced"`
}

type UpdateRecipeList struct {
	ID	  int `json:"ID"`
	Synced bool    `json:"synced"`
}