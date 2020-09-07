package main

import (
	"context"

	_ "github.com/go-kivik/couchdb/v4"
	kivik "github.com/go-kivik/kivik/v4"
)

var dbClient *kivik.DB
var connectString string

func setDBCredentials ( username string, password string , address string )  { 
	connectString = "http://" + username + ":" + password + "@" + address + ":5984/"
}

func connectDB( dbName string ) *kivik.DB {
	client, err := kivik.New("couch", connectString ) 

	if( err != nil ) {
		panic( err )
	}

	exists, _ := client.DBExists( context.TODO(), dbName )

	if( !exists ) {
		err := client.CreateDB( context.TODO(), dbName, nil )

		if( err != nil ) {
			panic( err )
		}
	
	}

	return client.DB(dbName) 
}

func getDB( dbName string ) *kivik.DB {
	return connectDB( dbName )
}
