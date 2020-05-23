	package main

import ( 
	"fmt"
)

var locationValue = map[string]int{
	"PassThrough": 0,
	"Mash": 1,
	"Adjunct1": 2,
	"Adjunct2": 3,
	"Adjunct3": 4,
	"Adjunct4": 5,
	"Pause": 6,
}

func createRecipeControllerRespMsg( ID int ) RecipeControllerRespMsg {

	var recipeControllerMsg RecipeControllerRespMsg

	if( len(  xmlRecipes) < ID ) {
		panic( fmt.Errorf("Recipe not found: %d", ID ) )
	}

	recipe := xmlRecipes[ID]

	recipeControllerMsg.ID = ID 
	recipeControllerMsg.Name = setString( recipe.Name, 19 )
	recipeControllerMsg.StartWater = recipe.Waters[0].Water.Amount
	

	for _, step := range recipe.Zymatic.Steps {
		recipeControllerMsg.Steps = append( recipeControllerMsg.Steps, convertSteps( step ) )
	}

	recipeControllerMsg.TypeCode = "Beer"
	
	return recipeControllerMsg
}

func setString( val2 string, maxSize int ) string {
	if ( len( val2 ) > maxSize ) {
		return val2[:20]
	}

	return val2
}
 
func convertSteps( xmlStep XmlZymaticStep ) RecipeControllerStepMsg {
	var stepMsg RecipeControllerStepMsg
	stepMsg.Drain = xmlStep.Drain
	stepMsg.Location = locationValue[xmlStep.Location]
	stepMsg.Name = setString( xmlStep.Name,19 )
	stepMsg.Temp = toFahrenheit( xmlStep.Temp )
	stepMsg.Time = xmlStep.Time
	return stepMsg
}

func toCelcius( fahrenheit int ) int {
	return ( (fahrenheit - 32 ) * 5 ) / 9
}

func toFahrenheit( celcius int ) int {
	return celcius * 9 / 5 + 32
}

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
