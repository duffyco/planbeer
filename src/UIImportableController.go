package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

var ImportableRecipes []ImportableRecipe
var ImportableXmlRecipes []XmlRecipe
var ImportableSessions []ImportableSession
var ImportableSessionCSVs []SessionCSVFileFormat

var	importRecipePath = getConfiguration().RecipePath
var importSessionPath = getConfiguration().SessionPath

func DeleteImportableSession( id int ) DefaultRespMsg {
	return DeleteImportableItem( id, importSessionPath )
}

func DeleteImportableRecipe( id int ) DefaultRespMsg {
	return DeleteImportableItem( id, importRecipePath )
}

func DeleteImportableItem( id int, importPath string ) DefaultRespMsg {
		var respMsg DefaultRespMsg

		files := ScanForRecipes(importPath)
		for i, xmls := range files {
			
			if( i == id ) {
				err := os.Remove( importPath + "/" + xmls.Name() )

				if( err == nil ) {
					respMsg.Value = "true";
					return respMsg;
				}

			}
	}	
		
	respMsg.Value = "false";
	return respMsg
}

func ListImportablesController() UIImportablesList {

	files := ScanForRecipes(importRecipePath)

	log.Printf( "Files: %d", len( files ) )

	var importableRecipes []ImportableRecipe
	var importableXmlRecipes []XmlRecipe

	//Load Recipes

	log.Printf( "importableXmlRecipes: %d ", len( importableXmlRecipes ) )
	for _, xmls := range files {
		xmlRecipeFile, err := RecipeFile(importRecipePath + "/" + xmls.Name())
		if( err != nil ) {
			continue;
		}
		
		log.Printf( "Adding: %s", xmlRecipeFile.Name )
		if !DBRecipeExists(xmlRecipeFile) {
			importableRecipes = append(importableRecipes, convertXmlRecipeToImportableRecipe(xmlRecipeFile))
			importableXmlRecipes = append(importableXmlRecipes, xmlRecipeFile)
		}
	}

	var importableSessionCSVs []SessionCSVFileFormat
	importableSessionCSVs = getImportableSessions(importSessionPath)

	var importableSessions []ImportableSession
	for _, importableSessionCSV := range importableSessionCSVs {
		lastItem := importableSessionCSV.Sessions[(len(importableSessionCSV.Sessions) - 1)]
		if !DBSessionExists(lastItem.ZSessionID) {
			importableSessions = append(importableSessions, convertSessionCSVFileToImportableSession(lastItem))
		}
	}

	log.Printf( "importableXmlRecipes[end]: %d", len( importableXmlRecipes ) );

	ImportableRecipes = importableRecipes
	ImportableXmlRecipes = importableXmlRecipes
	ImportableSessions = importableSessions
	ImportableSessionCSVs = importableSessionCSVs
	log.Printf( "ImportableXmlRecipes[end]: %d", len( ImportableXmlRecipes ) );

	return UIImportablesList{
		Sessions: importableSessions,
		Recipes:  importableRecipes,
	}
}

func UploadFilesController(inmsg UploadFilesData) bool {

	for i, fileData := range inmsg.Files {

		filename := inmsg.FileNames[i]
		var outputDirectory string

		if inmsg.Type == SessionUploadType {
			outputDirectory = importSessionPath
		} else if inmsg.Type == RecipeUploadType {
			outputDirectory = importRecipePath
		}

		err := ioutil.WriteFile(outputDirectory+"/"+filename, []byte(fileData) , 0644)
		
		if ( err != nil ) {
			panic( err );
		}


		return true
	}

	return false
}

func ImportSessionController(inmsg ImportSessionData) bool {
	for _, sess := range ImportableSessionCSVs {

		var firstEntry SessionCSVFormat
		for i, s := range sess.Sessions {
			if( s.ZSessionID != 0 ) {
				firstEntry = sess.Sessions[i]
				break;
			}
		}

		if firstEntry.ZSessionID == inmsg.SessionID {
			registerSession(getMachineByToken(inmsg.Machine), getRecipe(inmsg.RecipeID), firstEntry)
			for _, logEntry := range sess.Sessions {

				if( logEntry.ZSessionID == 0 ) {
					continue;
				}

				dateLayout := "1/2/2006 15:04:05 PM"
				t, err := time.Parse(dateLayout, logEntry.LogDate )

				if( err != nil ) {
					panic( err )
				}

				CreateSessionLogEntry(ConvertCSVtoLogMsg(logEntry), t.Unix())
			}
			return true
		}
	}
	return false
}

func ImportRecipeController(name string) bool {
	for i := range ImportableRecipes {
		if ImportableRecipes[i].Name == name {
			return UpdateRecipe(ImportableXmlRecipes[i])
		}
	}

	return false
}

type UIImportablesList struct {
	Sessions []ImportableSession `json:"ImportableSessions"`
	Recipes  []ImportableRecipe  `json:"ImportableRecipes"`
}

type ImportableRecipe struct {
	Name     string  `json:"Name"`
	Style    string  `json:"Style"`
	OG       float32 `json:"OG"`
	ABV      float32 `json:"ABV"`
	IBU      int     `json:"IBU"`
	SRM      int     `json:"SRM"`
	Date     string  `json:"Date"`
	BoilSize float32 `json:"BoilSize"`
}

type ImportableSession struct {
	ID      int    `json:"SessionID"`
	LogDate string `json:"LogDate"`
}

func convertSessionCSVFileToImportableSession(importCSV SessionCSVFormat) ImportableSession {
	return ImportableSession{
		ID:      importCSV.ZSessionID,
		LogDate: importCSV.LogDate,
	}
}

func convertXmlRecipeToImportableRecipe(importXML XmlRecipe) ImportableRecipe {
	return ImportableRecipe{
		Name:     importXML.Name,
		Style:    importXML.Style.Name,
		OG:       importXML.OG,
		ABV:      importXML.ABV,
		IBU:      importXML.IBU,
		SRM:      importXML.Color,
		Date:     importXML.Date,
		BoilSize: importXML.BoilSize,
	}
}

type ImportSessionData struct {
	SessionID int    `json:"SessionID"`
	Machine   string `json:"Machine"`
	RecipeID  int    `json:"RecipeID"`
}

type UploadFilesData struct {
	Type      string   `json:"Type"`
	Files     []string `json:Files`
	FileNames []string `json:FileNames`
}

const (
	SessionUploadType string = "Session"
	RecipeUploadType  string = "Recipe"
)
