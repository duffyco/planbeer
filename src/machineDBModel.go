package main

type MachineDBO struct {
	Token         string           `json:"MachineToken"`
	MachineState  ZStateResponse   `json:"ZState"`
	ID            string           `json:"_id"`
	MachineID     int              `json:"MachineID"`
	Rev           string           `json:"_rev,omitempty"`
	CurrentStatus MachineDBOStatus `json:"CurrentStatus"`
}

type MachineDBOStatus struct {
	LastSeen      string `json:"LastSeen"`
	Status        string `json:"Status"`
	SessionID     int    `json:"SessionID"`
	SessionType   int    `json:"SessionType"`
	ErrorCode     int    `json:"ErrorCode"`
	PauseReason   int    `json:"PauseReason"`
	TimeRemaining int    `json:"TimeRemaining"`
}