package main

import "time"

func GetLogsForSession(sessionID int) UISessionData {
	var sessionLog UISessionLog

	sessionLogs := GetSessionLogs(sessionID)
	session := GetSession(sessionID)

	for _, logs := range sessionLogs {
		sessionLog.Name = session.SessionMsg.Name
		sessionLog.ThermoBlockTemp = append(sessionLog.ThermoBlockTemp, logs.ThermoBlockTemp)
		sessionLog.WortTemp = append(sessionLog.WortTemp, logs.WortTemp)
		sessionLog.AmbientTemp = append(sessionLog.AmbientTemp, logs.AmbientTemp)
		sessionLog.DrainTemp = append(sessionLog.DrainTemp, logs.DrainTemp)
		sessionLog.TargetTemp = append(sessionLog.TargetTemp, logs.TargetTemp)
		sessionLog.StepName = append(sessionLog.StepName, logs.StepName)
		sessionLog.ValvePosition = append(sessionLog.ValvePosition, logs.ValvePosition)
		sessionLog.KegPumpOn = append(sessionLog.KegPumpOn, logs.KegPumpOn)
		sessionLog.DrainPumpOn = append(sessionLog.DrainPumpOn, logs.DrainPumpOn)
		sessionLog.PauseReason = append(sessionLog.PauseReason, logs.PauseReason)
		//sessionLog.LogDate = append(sessionLog.LogDate, time.Unix(logs.LogDate, 0).Format("2006-01-02T15:04:05.000"))
		sessionLog.LogDate = append(sessionLog.LogDate, time.Unix(logs.LogDate, 0).Format("15:04:05"))
	}

	return UISessionData{
		LogData: sessionLog,
	}
}

type UISessionData struct {
	LogData UISessionLog `json:"LogData"`
}

type UISessionLog struct {
	Name            string    `json:"Name"`
	ThermoBlockTemp []float32 `json:"ThermoBlockTemp"`
	WortTemp        []float32 `json:"WortTemp"`
	AmbientTemp     []float32 `json:"AmbientTemp"`
	DrainTemp       []float32 `json:"DrainTemp"`
	TargetTemp      []int     `json:"TargetTemp"`
	StepName        []string  `json:"StepName"`
	ValvePosition   []int     `json:"ValvePosition"`
	KegPumpOn       []int     `json:"KegPumpOn"`
	DrainPumpOn     []int     `json:"DrainPumpOn"`
	PauseReason     []int     `json:"PauseReason"`
	LogDate         []string  `json:"LogDate"`
}