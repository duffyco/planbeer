package main

import (
	"time"
) 

var machineList []Machine

//@TODO: Machines are assigned an ID.  Permantently for some reason.  Should be token only.
// UI compensates for this.

func getMachines() []Machine {
	loadMachines();

	for _, mach := range machineList {
		lastSeen, _ := time.Parse( "2006-01-02T15:04:05.000Z07:00", mach.CurrentStatus.LastSeen )

		today := time.Now().UTC()
		machineActiveWindow := lastSeen.Add( 2 * time.Minute )
		
		if(  today.After( machineActiveWindow) ) {
			updateMachineStatusIdle( mach.Token )
		}
	}

	return machineList
}

func updateMachineStatusIdle( token string ) {
	currentMachine := getMachineByToken( token )
	currentMachine.CurrentStatus.ErrorCode = 0
	currentMachine.CurrentStatus.PauseReason = 0
	currentMachine.CurrentStatus.SessionID = 0
	currentMachine.CurrentStatus.SessionType = 0
	currentMachine.CurrentStatus.TimeRemaining = 0
	currentMachine.CurrentStatus.Status = "Idle"
	updateMachineStatus( currentMachine )
}

func updateActiveSession( token string, sessionType int, sessionID int, timeleft int ) {
	currentMachine := getMachineByToken( token )
	currentMachine.CurrentStatus.SessionID = sessionID
	currentMachine.CurrentStatus.SessionType = sessionType
	currentMachine.CurrentStatus.TimeRemaining = timeleft
	currentMachine.CurrentStatus.Status = "Starting"
	updateMachineStatus( currentMachine )
}

func updateMachineBrewingStatus( token string, sessionID int, timeleft int, errorCode int, pauseReason int ) { 
	currentMachine := getMachineByToken( token )
	currentMachine.CurrentStatus.SessionID = sessionID
	currentMachine.CurrentStatus.TimeRemaining = timeleft
	currentMachine.CurrentStatus.ErrorCode = errorCode
	currentMachine.CurrentStatus.PauseReason = pauseReason
	currentMachine.CurrentStatus.Status = "Brewing"
	updateMachineStatus( currentMachine )
}


func updateMachineStatus( machine Machine ) {
	DBUpdateMachineStatus( machine )
}

func getMachineByToken(token string) Machine {
	for _, mach := range machineList {
		if mach.Token == token {
			return mach
		}
	}

	return Machine{}
}

func deleteMachine(token string) bool {
	var retVal bool = DBDeleteMachine(token)
	loadMachines()
	return retVal
}

func getMachine(mli int) Machine {
	return machineList[mli]
}

func loadMachines() {
	var dboms []MachineDBO
	dboms = DBGetMachines("machine")

	var maxID int = -1
	for _, mach := range dboms {
		if mach.MachineID > maxID {
			maxID = mach.MachineID
		}
	}

	machineList = make([]Machine, maxID+1)

	for _, machine := range dboms {
		machineList[machine.MachineID] = convertDBMachineToMachine(machine)
	}
}

func updateMachine(token string, resp ZStateResponse) ZStateResponse {
	foundRecipe := DBUpdateMachine("machine", token, len(machineList), resp)
	loadMachines()

	return foundRecipe.MachineState
}
