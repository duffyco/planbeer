package main

type ZState struct {
	CurrentFirmware string `json:"CurrentFirmware"`
	BoilerType      int    `json:"BoilerType"`
}

type ZStateStats struct {
	DirtySessionsSinceClean int `json:"DirtySessionsSinceClean"`
	LastSessionType         int `json:"LastSessionType"`
	ResumableSessionID      int `json:"ResumableSessionID"`
}

type ZStateResponse struct {
	IsUpdated         bool        `json:"IsUpdated"`
	IsRegistered      bool        `json:"IsRegistered"`
	TokenExpired      bool          `json:"TokenExpired"`
	UpdateAddress     string        `json:"UpdateAddress"`
	RegistrationToken string        `json:"RegistrationToken"`
	BoilerType        int           `json:"BoilerType"`
	CurrentFirmware   string        `json:"CurrentFirmware"`
	UpdateToFirmware  *string        `json:"UpdateToFirmware"`
	ProgramUri        *string        `json:"ProgramUri"`
	Alias             string        `json:"Alias"`
	SessionStats      ZStateStats `json:"SessionStats"`
	ZBackendError     int           `json:"ZBackendError"`
}

func createZStateResponse( resp ZStateResponse ) ZStateResponse {
	
	var respStats ZStateStats

	respStats.DirtySessionsSinceClean = 2
	respStats.LastSessionType = 0
	respStats.ResumableSessionID = -1
	
	resp.IsUpdated = true
	resp.IsRegistered = true
	resp.TokenExpired = false
	resp.UpdateAddress = "-1"
	resp.RegistrationToken = "-1"
	resp.BoilerType = 1
	resp.CurrentFirmware = "0.0.116"
	resp.UpdateToFirmware = nil
	resp.ProgramUri = nil
	resp.Alias = "Z2"
	resp.SessionStats = respStats
	resp.ZBackendError = 0

	return resp
}
