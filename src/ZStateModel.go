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

