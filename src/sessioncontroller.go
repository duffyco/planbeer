package main

import (
	"time"

	guuid "github.com/google/uuid"
)

var sessions = map[int]SessionHandle{}

type SessionHandle struct {
	sessionMsg SessionControllerMsg
	sessionRecipeID int
	recipeDescription XmlRecipe
}


var sessionType = map[string]int{
	"Rinse": 0,
	"Clean": 1,
	"Drain": 2,
	"Rack Beer": 3,
	"Circulate": 4,
	"Sous Vide": 5,
	"Beer": 6,
	"All Grain": 6,
	"Coffee": 12,
	"Chill": 13,
}

var zProgramID = map[string]int{
	"Rinse": 1,
	"Drain": 2,
	"Rack Beer": 3,
	"Circulate": 4,
	"Sous Vide": 6,
	"Clean": 12,
	"Beer/Coffee": 24,
	"All Grain": 24,
	"Chill": 27,
}

func GetSessionTypeFromRecipe( recipe Recipe ) int {
	return sessionType[recipe.XmlRecipe.Type];
}

func GetProgramIDFromRecipe( recipe Recipe ) int {
	return zProgramID[recipe.XmlRecipe.Type];
} 

func createSessionControllerRespMsg( inMsg SessionControllerMsg, token string ) SessionControllerRespMsg {
	var outMsg SessionControllerRespMsg
	curSessionID := CreateSession( inMsg, token )

	updateActiveSession( token, inMsg.SessionType, curSessionID, inMsg.DurationSec)

	outMsg.ID = curSessionID
	outMsg.ZSeriesID = 1335
	outMsg.ProfileID = 99999
	outMsg.SessionType = inMsg.SessionType
	outMsg.ZProgramId = inMsg.ZProgramId
	outMsg.Name = inMsg.Name
	
	outMsg.RecipeID = inMsg.RecipeID

	outMsg.FirmwareVersion = inMsg.FirmwareVersion
	outMsg.DurationSec = inMsg.DurationSec
	outMsg.MaxTemp = inMsg.MaxTemp
	outMsg.GroupSession = false
	outMsg.CreationDate = time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	//This works
	//outMsg.CreationDate = "2020-05-05T20:20:28.503"
	outMsg.GUID = guuid.New().String()
	outMsg.Deleted = false
	outMsg.Lat = 59.565833
	outMsg.Lng = -108.614444
	outMsg.CityLat = outMsg.Lat
	outMsg.CityLng = outMsg.Lng
	outMsg.Active = true
	outMsg.SessionLogs = []string {}

	// Create SessionControllerRespMsg

	return outMsg
}