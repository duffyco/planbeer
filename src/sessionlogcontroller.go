package main

import (
	"strconv"
	"time"
)

func createSessionLogRespMsg( inMsg SessionLogMsg, token string ) SessionLogRespMsg {
	var outMsg SessionLogRespMsg

	sessObj := GetSession( inMsg.ZSessionID )

	CreateSessionLogEntry( 	inMsg, time.Now().Unix() );

	updateMachineBrewingStatus( token, inMsg.ZSessionID, inMsg.SecondsRemaining, inMsg.ErrorCode, inMsg.PauseReason )

	outMsg.ID = int( time.Now().Unix() )
	outMsg.ZSessionID, _ = strconv.Atoi( sessObj.ID )
	outMsg.LogDate = time.Now().Format("2006-01-02T15:04:05.000Z07:00")
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