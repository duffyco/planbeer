package main

import (
	"log"
)


func SessionExists(id int) bool {
	return DBSessionExists(id)
}

func GetSessions() []Session {
	log.Printf("GetSessions()")
	var retSessions []Session 

	sessDBOs := DBGetSessions( "sessions" )

	for _, sessDBO := range sessDBOs {
		retSessions = append( retSessions, convertDBSessionToSession( sessDBO ) );
	}

	log.Printf("GetSessions-end()")
	return retSessions;
}

func GetSession(id int) Session {

	if DBSessionExists(id) {
		var foundSess SessionDBO
		foundSess, _ = DBFindSession(id)
		return convertDBSessionToSession( foundSess )
	}

	//@TODO: Is this a good idea??
	return Session{}
}

func DeleteSession( id int ) bool {
	return DBDeleteSession(id);
}

func importSessions(pathname string) {
	importSess := getImportableSessions(pathname)
	log.Printf("Importable Sessions: %d", len( importSess) )

	machine := getMachines()[0]
	recipe := getRecipe(0)


	for _, file := range importSess {
		registerSession( machine, recipe, file.Sessions[0] );
		for _, sess := range file.Sessions {
			CreateSessionLogEntry( ConvertCSVtoLogMsg( sess ), int64( sess.ID ) );
		}	
	}
}

func GetSessionLogs( sessionID int ) []SessionLog {
	return DBGetSessionLogs( sessionID )
}


func CreateSession( inSessionMsg SessionControllerMsg, token string ) int {
	return DBCreateSession( inSessionMsg, token );
}

func CreateSessionWithID( inSessionMsg SessionControllerMsg, token string, id int ) int {
	return DBCreateSessionWithID( inSessionMsg, token, id );
}

func registerSession( machine Machine, recipe Recipe, sess SessionCSVFormat ) {
	sessParams:= &SessionProgramParams {
		Abv: -1,
		Ibu: -1,
		Duration: 0,
		Water: recipe.XmlRecipe.BoilSize,
		Intensity: 0,
		Temperature: 0,
	}
	
	sessionMsg := &SessionControllerMsg {
		SessionType: GetSessionTypeFromRecipe( recipe ),
		RecipeID: recipe.RecipeID,
		Name: recipe.XmlRecipe.Name, 
		FirmwareVersion: machine.MachineState.CurrentFirmware,
		DurationSec: CalculateDurationFromRecipe( recipe ),
		PressurePa: 100703.8,
		MaxTemp: float32( recipe.XmlRecipe.Zymatic.BoilTemp ),
		MaxTempAddedSec: 0,
		ZProgramId: GetProgramIDFromRecipe( recipe ),
		GroupSession: true,
		ProgramParams: *sessParams,
	}

	sessID := CreateSessionWithID( *sessionMsg, machine.Token, sess.ZSessionID )

	log.Printf( "Created Session at %d", sessID )

}

func CreateSessionLogEntry( sessLogMsg SessionLogMsg, logDate int64 ) {
	DBCreateSessionLogEntry( SessionLog {
		SessionLogMsg: sessLogMsg,
		LogDate: logDate,
	}  );
}

func ConvertCSVtoLogMsg( csvEntry SessionCSVFormat ) SessionLogMsg {
	
	var sessLogMsg SessionLogMsg
	sessLogMsg.ZSessionID = csvEntry.ZSessionID
	sessLogMsg.ThermoBlockTemp = csvEntry.ThermoBlockTemp
	sessLogMsg.WortTemp = csvEntry.WortTemp
	sessLogMsg.AmbientTemp = csvEntry.AmbientTemp
	sessLogMsg.DrainTemp = csvEntry.DrainTemp
	sessLogMsg.TargetTemp = int( csvEntry.TargetTemp )
	sessLogMsg.StepName = csvEntry.StepName
	sessLogMsg.ValvePosition = csvEntry.ValvePosition
	if( csvEntry.KegPumpOn ) {
		sessLogMsg.KegPumpOn = 1 
	} else {
		sessLogMsg.KegPumpOn = 0
	}

	if( csvEntry.DrainPumpOn ) {
		sessLogMsg.DrainPumpOn = 1 
	} else {
		sessLogMsg.DrainPumpOn = 0
	}
	sessLogMsg.PauseReason = csvEntry.PauseReason
	sessLogMsg.ErrorCode = csvEntry.ErrorCode
	sessLogMsg.Rssi = csvEntry.Rssi
	sessLogMsg.NetSend = csvEntry.NetSend
	sessLogMsg.NetWait = csvEntry.NetWait
	sessLogMsg.NetRecv = csvEntry.NetRecv 
	sessLogMsg.SecondsRemaining = 0

	return sessLogMsg
}