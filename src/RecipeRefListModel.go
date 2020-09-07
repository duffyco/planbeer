package main

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
