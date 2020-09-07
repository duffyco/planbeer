package main

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-kivik/couchdb/v4"
	guuid "github.com/google/uuid"
)

func DBGetMachines( databaseName string ) []MachineDBO  {

	var mlDB []MachineDBO;
			
	rows, err := getDB("machine").AllDocs( context.TODO(), map[string]interface{}{"include_docs": true} )

	if( err != nil ) {
		panic( err )
	}

	for rows.Next() {
		var sdbo MachineDBO;
		if err := rows.ScanDoc(&sdbo); err != nil {
			panic(err)
		}

		mlDB = append( mlDB, sdbo )
	}

	return mlDB;
}

func DBDeleteMachine( token string ) bool {

	row := getDB( "machine").Get( context.TODO(), token )

	var dbor MachineDBO
	err := row.ScanDoc(&dbor); 

	if( err != nil ) {
		return false
	}

	_, err = getDB( "machine").Delete( context.TODO(), token, row.Rev )
		
	if( err != nil ){
		return false;
	}

	return true;
}

func DBUpdateMachineStatus( newStatus Machine ) {
	row := getDB( "machine").Get( context.TODO(), newStatus.Token )

	var dbor MachineDBO
	err := row.ScanDoc(&dbor); 

	curStatus := newStatus.CurrentStatus

	if( err == nil ) {
		curStatus.LastSeen = time.Now().Local().Format("2006-01-02T15:04:05.000Z07:00") 
		dbor.CurrentStatus = MachineDBOStatus(curStatus)
		rev, err := getDB("machine").Put( context.TODO(), newStatus.Token, dbor )
	
		if( err != nil ) {
			panic( err )
		}
		
		fmt.Printf("Machine Inserted with revision %s\n", rev)
	}
}


func DBUpdateMachine( databaseName string, token string, machineID int, store ZStateResponse ) MachineDBO {
	dbo := &MachineDBO{ 
		Token: token, 
		MachineState: store, 
		MachineID: machineID,
		ID: guuid.New().String(),
	}

	row := getDB( "machine").Get( context.TODO(), token )

	var dbor MachineDBO
	err := row.ScanDoc(&dbor); 

	if( err != nil ) {
		rev, err := getDB("machine").Put( context.TODO(), token, dbo )
	
		if( err != nil ) {
			panic( err )
		}

		fmt.Printf("%s inserted with revision %s\n", databaseName, rev)
		return *dbo
	}

	return dbor
}
