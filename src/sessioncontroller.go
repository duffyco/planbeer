package main

import (
	"math/rand"
)

var sessions = map[int]SessionHandle{}

var sessionLog = map[int][]string {}

type SessionHandle struct {
	sessionMsg SessionControllerMsg
	sessionRecipeID int
	recipeDescription XmlRecipe
}

func NewSessionHandle( inSessionMsg SessionControllerMsg ) SessionHandle {
	h := SessionHandle { sessionMsg: inSessionMsg }
	h.sessionRecipeID = inSessionMsg.RecipeID
	// @TODO: What if it doesn't find the recipe?
	if( sessionType[inSessionMsg.SessionType] == "Beer" ) {
		h.recipeDescription = xmlRecipes[inSessionMsg.RecipeID]
	}

	return h
}

var sessionType = map[int]string{
	0: "Rinse",
	1: "Clean",
	2: "Drain",
	3: "Rack Beer",
	4: "Circulate",
	5: "Sous Vide",
	6: "Beer",
	12: "Coffee",
	13: "Chill",
}

var zProgramID = map[int]string{
	1: "Rinse",
	2: "Drain",
	3: "Rack Beer",
	4: "Circulate",
	6: "Sous Vide",
	12: "Clean",
	24: "Beer/Coffee",
	27: "Chill",
}

func createSessionControllerRespMsg( inMsg SessionControllerMsg ) SessionControllerRespMsg {
	var outMsg SessionControllerRespMsg
	curSessionID := rand.Intn( 99999 )
	sessions[curSessionID] = NewSessionHandle( inMsg )
	sessionLog[curSessionID] = []string{}
//	curSession := sessions[curSessionID]

	outMsg.ID = curSessionID
	outMsg.ZSeriesID = 1335
	outMsg.ProfileID = 33250
	outMsg.SessionType = inMsg.SessionType
	outMsg.ZProgramId = inMsg.ZProgramId
	outMsg.Name = inMsg.Name
	
	outMsg.RecipeID = inMsg.RecipeID

	outMsg.FirmwareVersion = inMsg.FirmwareVersion
	outMsg.DurationSec = inMsg.DurationSec
	outMsg.MaxTemp = inMsg.MaxTemp
	outMsg.GroupSession = false
	outMsg.CreationDate = "2020-05-05T20:20:28.503"
	outMsg.GUID = "9D574062468E4783A0BD5BB3D853AC02"
	outMsg.Deleted = false
	outMsg.Lat = 45.421528
	outMsg.Lng = -75.697189
	outMsg.CityLat = outMsg.Lat
	outMsg.CityLng = outMsg.Lng
	outMsg.Active = true
	outMsg.SessionLogs = sessionLog[curSessionID] 

	// Create SessionControllerRespMsg

	return outMsg
}


type SessionControllerMsg struct {
	SessionType  int `json:"SessionType"`
	RecipeID  int `json:"RecipeID"`
	Name  string `json:"Name"` 
	FirmwareVersion  string `json:"FirmwareVersion"`
	DurationSec  int `json:"DurationSec"`
	PressurePa  float32 `json:"PressurePa"`
	MaxTemp  float32 `json:"MaxTemp"`
	MaxTempAddedSec  int `json:"MaxTempAddedSec"`
	ZProgramId  int `json:"ZProgramId"`
	GroupSession  bool `json:"GroupSession"`
	ProgramParams SessionProgramParams `json:"ProgramParams"`
}

type SessionProgramParams struct {
	Abv  int `json:"Abv"`
	Ibu  int `json:"Ibu"`
	Duration  int `json:"Duration"`
	Water  float32 `json:"Water"`
	Intensity  int `json:"Intensity"`
	Temperature  int `json:"Temperature"`
}

type SessionControllerRespMsg struct {
	ID  int `json:"ID"`
	ZSeriesID  int `json:"ZSeriesID"`
	ProfileID  int `json:"ProfileID"`
	SessionType  int `json:"SessionType"`
	ZProgramId  int `json:"ZProgramId"`
	LastLogID *string `json:"LastLogID"`
	Name      string    `json:"Name"`
	RecipeID  int `json:"RecipeID"`
	FirmwareVersion  string `json:"FirmwareVersion"`
	StillUID *string `json:"StillUID"`
	StillVer *string `json:"StillVer"`
	GroupSession  bool `json:"GroupSession"`
	RecipeGuid *string `json:"RecipeGuid"`
	CreationDate      string    `json:"CreationDate"`
	ClosingDate      *string    `json:"ClosingDate"`
	GUID      string    `json:"GUID"`
	Deleted  bool `json:"Deleted"`
	Notes      *string    `json:"Notes"`
	Lat  float32 `json:"Lat"`
	Lng  float32 `json:"Lng"`
	CityLat  float32 `json:"CityLat"`
	CityLng  float32 `json:"CityLng"`
	Active  bool `json:"Active"`
	DurationSec  int `json:"DurationSec"`
	Pressure  int `json:"Pressure"`
	MaxTemp  float32 `json:"MaxTemp"`
	MaxTempAddedSec  int `json:"MaxTempAddedSec"`
	ProgramParams SessionProgramParams `json:"ProgramParams"`
	SessionLogs []string `json:"SessionLogs"`
	SecondsRemaining      *string    `json:"SecondsRemaining"`
}
