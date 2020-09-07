package main

type RecipeDBO struct {
	ID        string    `json:"_id"`
	RecipeID  int       `json:"recipeID"`
	Rev       string    `json:"_rev,omitempty"`
	XmlRecipe XmlRecipe `json:"XmlRecipe"`
	Synced    bool      `json:"synced"`
}