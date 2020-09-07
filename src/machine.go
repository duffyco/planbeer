package main

type Machine struct {
	Token         string         `json:"MachineToken"`
	MachineState  ZStateResponse `json:"ZState"`
	ID            string         `json:"_id"`
	MachineID     int            `json:"MachineID"`
	CurrentStatus MachineStatus  `json:"CurrentStatus"`
}

type MachineStatus struct {
	LastSeen      string `json:"LastSeen"`
	Status        string `json:"Status"`
	SessionID     int    `json:"SessionID"`
	SessionType   int    `json:"SessionType"`
	ErrorCode     int    `json:"ErrorCode"`
	PauseReason   int    `json:"PauseReason"`
	TimeRemaining int    `json:"TimeRemaining"`
}

func convertDBMachineToMachine(dbo MachineDBO) Machine {
	var retMachine Machine

	retMachine.ID = dbo.ID
	retMachine.MachineID = dbo.MachineID
	retMachine.Token = dbo.Token
	retMachine.MachineState = dbo.MachineState
	retMachine.CurrentStatus.ErrorCode = dbo.CurrentStatus.ErrorCode
	retMachine.CurrentStatus.LastSeen = dbo.CurrentStatus.LastSeen
	retMachine.CurrentStatus.PauseReason = dbo.CurrentStatus.PauseReason
	retMachine.CurrentStatus.SessionID = dbo.CurrentStatus.SessionID
	retMachine.CurrentStatus.SessionType = dbo.CurrentStatus.SessionType
	retMachine.CurrentStatus.Status = dbo.CurrentStatus.Status
	retMachine.CurrentStatus.TimeRemaining = dbo.CurrentStatus.TimeRemaining

	return retMachine
}