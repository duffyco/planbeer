package main

type Recipe struct {
	ID int
	Name string
	Kind int
	Uri *string
	Abv int 
	Ibu int
}

func NewRecipe( inId int, inName string) Recipe {
	recipe := Recipe{}
	recipe.ID = inId
	recipe.Name = inName
	recipe.Kind = 0
	recipe.Uri = nil
	recipe.Abv = -1
	recipe.Ibu = -1

	return recipe
}

func GetRecipeDescription( rec Recipe ) RecipeDescription {
	var ret RecipeDescription;

	ret.Abv = rec.Abv
	ret.ID = rec.ID
	ret.Ibu = rec.Ibu
	ret.Kind = rec.Kind
	ret.Name = rec.Name
	ret.Uri = rec.Uri

	return ret;
}