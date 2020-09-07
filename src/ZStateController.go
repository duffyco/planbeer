package main

func createDefaultZStateResponse() ZStateResponse {
	var resp ZStateResponse
	var respStats ZStateStats

	respStats.DirtySessionsSinceClean = 0
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

func createZStateResponse(resp ZStateResponse, token string) ZStateResponse {
	return updateMachine(token, createDefaultZStateResponse())
}