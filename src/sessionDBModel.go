package main

type SessionDBO struct {
	ID              string               `json:"_id"`
	Rev             string               `json:"_rev,omitempty"`
	Creation        string               `json:"creationtime"`
	SessionMsg      SessionControllerMsg `json:"sessionmsg"`
	SessionRecipeID int                  `json:"recipeID"`
	XmlRecipe       XmlRecipe            `json:"XmlRecipe"`
	Token           string               `json:"token"`
}

//@TODO: Date of start?  How do I know if it's active?  Do I care?
