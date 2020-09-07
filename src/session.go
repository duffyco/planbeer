package main

type Session struct {
	ID              string               `json:"_id"`
	Creation        string               `json:"creationtime"`
	SessionMsg      SessionControllerMsg `json:"sessionmsg"`
	SessionRecipeID int                  `json:"recipeID"`
	XmlRecipe       XmlRecipe            `json:"XmlRecipe"`
	Token           string               `json:"Token"`
}

func convertDBSessionToSession(dbo SessionDBO) Session {
	var retSession Session

	retSession.Creation = dbo.Creation
	retSession.ID = dbo.ID
	retSession.SessionMsg = dbo.SessionMsg
	retSession.SessionRecipeID = dbo.SessionRecipeID
	retSession.XmlRecipe = dbo.XmlRecipe
	retSession.Token = dbo.Token

	return retSession
}