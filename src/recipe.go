package main

type Recipe struct {
	ID        string    `json:"_id"`
	RecipeID  int       `json:"recipeID"`
	XmlRecipe XmlRecipe `json:"XmlRecipe"`
	Synced    bool      `json:"synced"`
}

func convertDBRecipeToRecipe(dbo RecipeDBO) Recipe {
	var retRecipe Recipe

	retRecipe.RecipeID = dbo.RecipeID
	retRecipe.ID = dbo.ID
	retRecipe.XmlRecipe = dbo.XmlRecipe
	retRecipe.Synced = dbo.Synced

	return retRecipe
}

func convertRecipeToDBRecipe(rec Recipe) RecipeDBO {
	var retRecipe RecipeDBO

	retRecipe.RecipeID = rec.RecipeID
	retRecipe.ID = rec.ID
	retRecipe.XmlRecipe = rec.XmlRecipe
	retRecipe.Synced = rec.Synced

	return retRecipe
}

func CalculateDurationFromRecipe(recipe Recipe) int {
	return 999999
}
