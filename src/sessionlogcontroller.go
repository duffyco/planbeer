package main

import (
	"log" 
	"encoding/json"
)

func createSessionLogRespMsg( inMsg SessionLogMsg ) SessionLogRespMsg {
	var outMsg SessionLogRespMsg

	curSessionID := inMsg.ZSessionID
	curLogSessionID := 11000000
	out, err := json.Marshal( &inMsg )

	if err != nil {
        panic (err)
    }

	sessionLog[curLogSessionID] = append( sessionLog[curLogSessionID], string( out ) )

	log.Printf( string( out ) )

	// @TODO: HELP!  ID comes from where
	outMsg.ID = curLogSessionID
	outMsg.ZSessionID = curSessionID
	
	outMsg.LogDate = "2020-05-05T20:21:02.46"
	outMsg.ThermoBlockTemp = inMsg.ThermoBlockTemp
	outMsg.WortTemp = inMsg.WortTemp
	outMsg.AmbientTemp = inMsg.AmbientTemp
	outMsg.DrainTemp = inMsg.DrainTemp
	outMsg.TargetTemp = inMsg.TargetTemp
	outMsg.ValvePosition = inMsg.ValvePosition
	outMsg.KegPumpOn = !( inMsg.KegPumpOn == 0 )
	outMsg.DrainPumpOn = !( inMsg.DrainPumpOn == 0 )
	outMsg.StepName = inMsg.StepName
	outMsg.ErrorCode = inMsg.ErrorCode
	// @TODO: Something here barfs on the line below.  Not sure why.
	// Maybe this is the query: {"ZSessionID":48084,"ThermoBlockTemp":19.799597,"WortTemp":19.851448,"AmbientTemp":23.220558,"DrainTemp":20.419252,"TargetTemp":0,"StepName":"Preparing To Brew","ValvePosition":2,"KegPumpOn":1,"DrainPumpOn":1,"PauseReason":0,"ErrorCode":0,"rssi":-41,"netSend":2917,"netWait":2966,"netRecv":120,"SecondsRemaining":21460}
	// This one is fine: 0 {"ZSessionID":48084,"ThermoBlockTemp":19.926075,"WortTemp":20.28869,"AmbientTemp":20.687141,"DrainTemp":18.366287,"TargetTemp":0,"StepName":"Preparing To Brew","ValvePosition":2,"KegPumpOn":1,"DrainPumpOn":1,"PauseReason":0,"ErrorCode":0,"rssi":-31,"netSend":2917,"netWait":3490,"netRecv":87,"SecondsRemaining":21459}
	outMsg.PauseReason = inMsg.PauseReason
	outMsg.Rssi = inMsg.Rssi
	outMsg.NetSend = inMsg.NetSend
	outMsg.NetRecv = inMsg.NetRecv
	outMsg.NetWait = inMsg.NetWait
	return outMsg
}


type SessionLogMsg struct {
	ZSessionID  int `json:"ZSessionID"`
	ThermoBlockTemp  float32 `json:"ThermoBlockTemp"`
	WortTemp  float32 `json:"WortTemp"`
	AmbientTemp  float32 `json:"AmbientTemp"`
	DrainTemp  float32 `json:"DrainTemp"`
	TargetTemp  float32 `json:"TargetTemp"`
	StepName      string    `json:"StepName"`
	ValvePosition  float32 `json:"ValvePosition"`
	KegPumpOn  int `json:"KegPumpOn"`
	DrainPumpOn  int `json:"DrainPumpOn"`
	PauseReason  int `json:"PauseReason"`
	ErrorCode  int `json:"ErrorCode"`
	Rssi  int `json:"rssi"`
	NetSend  int `json:"netSend"`
	NetWait  int `json:"netWait"`
	NetRecv  int `json:"netRecv"` 
	SecondsRemaining  int `json:"SecondsRemaining"`
}

type SessionLogRespMsg struct {
	ID  int `json:"ID"`
	ZSessionID  int `json:"ZSessionID"`
	LogDate  string `json:"LogDate"`
	ThermoBlockTemp  float32 `json:"ThermoBlockTemp"`
	WortTemp  float32 `json:"WortTemp"`
	AmbientTemp  float32 `json:"AmbientTemp"`
	DrainTemp  float32 `json:"DrainTemp"`
	TargetTemp  float32 `json:"TargetTemp"`
	ValvePosition  float32 `json:"ValvePosition"`
	KegPumpOn  bool `json:"KegPumpOn"`
	DrainPumpOn  bool `json:"DrainPumpOn"`
	StepName      string    `json:"StepName"`
	ErrorCode  int `json:"ErrorCode"`
	PauseReason  int `json:"PauseReason"`
	Rssi  int `json:"rssi"`
	NetSend  int `json:"netSend"`
	NetWait  int `json:"netWait"`
	NetRecv  int `json:"netRecv"`
	SecondsRemaining  *string `json:"SecondsRemaining"`
	StillSessionLog  *string `json:"StillSessionLog"`
	StillSessionLogID  *string `json:"StillSessionLogID"`
}
