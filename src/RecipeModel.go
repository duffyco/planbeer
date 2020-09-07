package main

type RecipeControllerStepMsg struct {
	Name  string `json:"Name"`
	Temp  int `json:"Temp"`
	Time  int `json:"Time"` 
	Location  int `json:"Location"`
	Drain  int `json:"Drain"`
}

type RecipeControllerRespMsg struct {
	ID  int `json:"ID"`
	Name      string    `json:"Name"`
	StartWater  float32  `json:"StartWater"`
	TypeCode  string `json:"TypeCode"`
	Steps []RecipeControllerStepMsg `json:"Steps"`
}