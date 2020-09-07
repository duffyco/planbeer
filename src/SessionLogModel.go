package main

type SessionLogMsg struct {
	ZSessionID       int     `json:"ZSessionID"`
	ThermoBlockTemp  float32 `json:"ThermoBlockTemp"`
	WortTemp         float32 `json:"WortTemp"`
	AmbientTemp      float32 `json:"AmbientTemp"`
	DrainTemp        float32 `json:"DrainTemp"`
	TargetTemp       int     `json:"TargetTemp"`
	StepName         string  `json:"StepName"`
	ValvePosition    int     `json:"ValvePosition"`
	KegPumpOn        int     `json:"KegPumpOn"`
	DrainPumpOn      int     `json:"DrainPumpOn"`
	PauseReason      int     `json:"PauseReason"`
	ErrorCode        int     `json:"ErrorCode"`
	Rssi             int     `json:"rssi"`
	NetSend          int     `json:"netSend"`
	NetWait          int     `json:"netWait"`
	NetRecv          int     `json:"netRecv"`
	SecondsRemaining int     `json:"SecondsRemaining"`
}

type SessionLog struct {
	SessionLogMsg
	LogDate int64 `json:"LogDate"`
}

//Changed ValvePosition / TargetTemp from float32 to int

type SessionLogRespMsg struct {
	ID                int     `json:"ID"`
	ZSessionID        int     `json:"ZSessionID"`
	LogDate           string  `json:"LogDate"`
	ThermoBlockTemp   float32 `json:"ThermoBlockTemp"`
	WortTemp          float32 `json:"WortTemp"`
	AmbientTemp       float32 `json:"AmbientTemp"`
	DrainTemp         float32 `json:"DrainTemp"`
	TargetTemp        int     `json:"TargetTemp"`
	ValvePosition     int     `json:"ValvePosition"`
	KegPumpOn         bool    `json:"KegPumpOn"`
	DrainPumpOn       bool    `json:"DrainPumpOn"`
	StepName          string  `json:"StepName"`
	ErrorCode         int     `json:"ErrorCode"`
	PauseReason       int     `json:"PauseReason"`
	Rssi              int     `json:"rssi"`
	NetSend           int     `json:"netSend"`
	NetWait           int     `json:"netWait"`
	NetRecv           int     `json:"netRecv"`
	SecondsRemaining  *string `json:"SecondsRemaining"`
	StillSessionLog   *string `json:"StillSessionLog"`
	StillSessionLogID *string `json:"StillSessionLogID"`
}
