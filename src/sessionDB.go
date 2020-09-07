package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-kivik/couchdb/v4"
)
func DBCreateSession( inSessionMsg SessionControllerMsg, token string) int {
	return DBCreateSessionWithID(inSessionMsg, token, rand.Intn( 99999 ) ); 
}

//@TODO: -1 comes in during exit???
func DBCreateSessionWithID( inSessionMsg SessionControllerMsg, token string, id int ) int {
	log.Printf( "DBCreateSessionWithID")

	sessDBO := &SessionDBO {
		ID: strconv.Itoa(id),
		Creation:  time.Now().Format("2006-01-02T15:04:05.000"),
		SessionMsg: inSessionMsg,
		SessionRecipeID: inSessionMsg.RecipeID,
		Token: token,
	}

	if( inSessionMsg.RecipeID > -1 ) {
		sessDBO.XmlRecipe = getRecipe(inSessionMsg.RecipeID).XmlRecipe
	}

	log.Printf( "DBCreateSessionWithID-Get")
	row := getDB( "sessions").Get( context.TODO(), sessDBO.ID )
	for row.ContentLength != 0 {
		row = getDB( "sessions").Get( context.TODO(), sessDBO.ID )
		sessDBO.ID = strconv.Itoa( rand.Intn( 99999 ) )
	}

	log.Printf( "DBCreateSessionWithID-Put")

	_, err := getDB("sessions").Put( context.TODO(), sessDBO.ID, sessDBO )
	
	//@TODO: Check if Session Exists??
	if( err != nil ) {
		//panic( err )
	}

	fmt.Printf("Started New Session (%s): %s....\n", sessDBO.ID, sessDBO.SessionMsg.Name )
	retValue, _ := strconv.Atoi( sessDBO.ID )

	log.Printf( "DBCreateSessionWithID-ends")
	return retValue
}

func DBSessionExists( idIn int ) bool {
	log.Printf( "DBSessionExists")
	row := getDB( "sessions").Get( context.TODO(), strconv.Itoa( idIn ) )
	log.Printf( "DBSessionExists-end")
	return ( row.ContentLength != 0 );
}

func DBFindSession( idIn int ) (SessionDBO, error) {

	log.Printf( "DBFindSession")
	row := getDB( "sessions").Get( context.TODO(), strconv.Itoa( idIn ) )
	
	if row.ContentLength == 0 {
		return SessionDBO{}, fmt.Errorf( "Session not found: %d", idIn)
	}

	log.Printf( "DBFindSession-Scandoc")
	var  sessObj SessionDBO
	err := row.ScanDoc(&sessObj);
	
	if( err != nil ){
		return SessionDBO{}, fmt.Errorf( "Cannot ScanDoc not found: %d", idIn)
	}

	log.Printf( "DBFindSession-ends")
	return sessObj, nil
}

func DBDeleteSession( idIn int ) bool {

	log.Printf( "DBDeleteSession")
	existDBO, _ := DBFindSession( idIn );

	_, err := getDB( "sessions").Delete( context.TODO(), strconv.Itoa( idIn ), existDBO.Rev )
		
	if( err != nil ){
		return false;
	}

	log.Printf( "DBDeleteSession-ends")
	return true;
}


func DBGetSessionLogs( id int ) []SessionLog{
	log.Printf( "DBGetSessionLogs")
	var logsList []SessionLog

	rows, err := getDB("sessionlog").AllDocs( context.TODO(), map[string]interface{}{"include_docs": true} )

	if( err != nil ) {
		panic( err )
	}

	log.Printf( "DBGetSessionLogs-ScanDoc")
	for rows.Next() {
		var slmdbo SessionLog
		if err := rows.ScanDoc(&slmdbo); err != nil {
			panic(err)
		}

		if( slmdbo.ZSessionID == id ) {
			logsList = append( logsList, slmdbo )
		}
	}

	log.Printf( "DBGetSessionLogs-end")
	return logsList;

}


func DBCreateSessionLogEntry( sessLogMsg SessionLog ) {
	//@TODO: I think the key needs to be [SessionID, Time].  This makes it unique + findable.
	
	log.Printf( "DBCreateSessionLogEntry")
	_, err := getDB( "sessionlog").Put( context.TODO(),  strconv.FormatInt( sessLogMsg.LogDate, 10), sessLogMsg )

	if( err != nil ){
		log.Printf( "Not adding record (already exists): %s", strconv.FormatInt( sessLogMsg.LogDate, 10))
	}
	log.Printf( "DBCreateSessionLogEntry-end")
}

func DBGetSessions( databaseName string ) []SessionDBO  {
	log.Printf( "DBGetSessions")
	var sessionList []SessionDBO
		
	log.Printf( "DBGetSessions-AllDocs")
	rows, err := getDB("sessions").AllDocs( context.TODO(), map[string]interface{}{"include_docs": true} )

	if( err != nil ) {
		panic( err )
	}

	for rows.Next() {
		log.Printf( "DBGetSessions-ScanDoc")		
		var sdbo SessionDBO
		if err := rows.ScanDoc(&sdbo); err != nil {
			panic(err)
		}

		sessionList = append( sessionList, sdbo )
	}

	log.Printf( "DBGetSessions-end")		
	return sessionList;
}