package main

import (
	"log"
	"strconv"
)

func ListSessionsController() UISessionList {

	log.Printf( "ListSessionsController")
	var sessionList UISessionList

	var dboSessionList []Session
	dboSessionList = GetSessions();

	for _, session := range dboSessionList {
		sessionList.Sessions = append(sessionList.Sessions, convertSessiontoSessionUIList(session))
	}

	return sessionList
}

func GetUISession(id int) UISession {
	return convertSessiontoSessionUIList( GetSession( id ) )
}

func DeleteUISession( id int ) DefaultRespMsg {
	var respMsg DefaultRespMsg

	respMsg.Value = strconv.FormatBool( DeleteSession( id ) );
	return respMsg
}


type UISessionList struct {
	Sessions []UISession `json:"Sessions"`
}

func convertSessiontoSessionUIList(inSession Session) UISession {
	var retSession UISession

	retSession.SessionID, _ = strconv.Atoi(inSession.ID)
	retSession.RecipeID = inSession.SessionRecipeID
	retSession.Token = inSession.Token
	retSession.BrewDate = inSession.Creation
	retSession.Type = inSession.SessionMsg.SessionType

	if( inSession.XmlRecipe.Type != ""  ) {
		retSession.Machine = inSession.XmlRecipe.Equipment.Name
		retSession.RecipeName = inSession.XmlRecipe.Name
		retSession.Style = inSession.XmlRecipe.Style.Name
		retSession.ABV = inSession.XmlRecipe.ABV
		retSession.IBU = inSession.XmlRecipe.IBU
		retSession.OG = inSession.XmlRecipe.OG
		retSession.SRM = inSession.XmlRecipe.Color
	} else {
		retSession.RecipeName = inSession.SessionMsg.Name
		retSession.Style = "Built In Program"
	}

	return retSession
}

type UISession struct {
	SessionID  int    `json:"ID"`
	Machine    string `json:"machinename"`
	RecipeName string `json:"recipename"`
	RecipeID   int    `json:"recipeid"`
	Token      string `json:"token"`
	Style      string `json:"style"`
	BrewDate   string `json:"brewdate"`
	ABV		   float32 `json:"ABV"`
	IBU		   int `json:"IBU"`
	OG		   float32 `json:"OG"`
	SRM		   int `json:"SRM"`
	Type		   int `json:"recipetype"`
}