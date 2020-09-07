package main

type SessionControllerMsg struct {
	SessionType  int `json:"SessionType"`
	RecipeID  int `json:"RecipeID"`
	Name  string `json:"Name"` 
	FirmwareVersion  string `json:"FirmwareVersion"`
	DurationSec  int `json:"DurationSec"`
	PressurePa  float32 `json:"PressurePa"`
	MaxTemp  float32 `json:"MaxTemp"`
	MaxTempAddedSec  int `json:"MaxTempAddedSec"`
	ZProgramId  int `json:"ZProgramId"`
	GroupSession  bool `json:"GroupSession"`
	ProgramParams SessionProgramParams `json:"ProgramParams"`
}

type SessionProgramParams struct {
	Abv  int `json:"Abv"`
	Ibu  int `json:"Ibu"`
	Duration  int `json:"Duration"`
	Water  float32 `json:"Water"`
	Intensity  int `json:"Intensity"`
	Temperature  int `json:"Temperature"`
}

type SessionControllerRespMsg struct {
	ID  int `json:"ID"`
	ZSeriesID  int `json:"ZSeriesID"`
	ProfileID  int `json:"ProfileID"`
	SessionType  int `json:"SessionType"`
	ZProgramId  int `json:"ZProgramId"`
	LastLogID *string `json:"LastLogID"`
	Name      string    `json:"Name"`
	RecipeID  int `json:"RecipeID"`
	FirmwareVersion  string `json:"FirmwareVersion"`
	StillUID *string `json:"StillUID"`
	StillVer *string `json:"StillVer"`
	GroupSession  bool `json:"GroupSession"`
	RecipeGuid *string `json:"RecipeGuid"`
	CreationDate      string    `json:"CreationDate"`
	ClosingDate      *string    `json:"ClosingDate"`
	GUID      string    `json:"GUID"`
	Deleted  bool `json:"Deleted"`
	Notes      *string    `json:"Notes"`
	Lat  float32 `json:"Lat"`
	Lng  float32 `json:"Lng"`
	CityLat  float32 `json:"CityLat"`
	CityLng  float32 `json:"CityLng"`
	Active  bool `json:"Active"`
	DurationSec  int `json:"DurationSec"`
	Pressure  int `json:"Pressure"`
	MaxTemp  float32 `json:"MaxTemp"`
	MaxTempAddedSec  int `json:"MaxTempAddedSec"`
	ProgramParams SessionProgramParams `json:"ProgramParams"`
	SessionLogs []string `json:"SessionLogs"`
	SecondsRemaining      *string    `json:"SecondsRemaining"`
}
