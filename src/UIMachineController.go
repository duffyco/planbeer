package main

import "strconv"

func DeleteUIMachine(token string) DefaultRespMsg {
	var respMsg DefaultRespMsg

	respMsg.Value = strconv.FormatBool(deleteMachine(token))
	return respMsg
}

func ListMachinesController() UIMachineList {

	var mlist UIMachineList

	var dboMachineList []Machine
	dboMachineList = getMachines()

	for _, machine := range dboMachineList {
		mlist.Machines = append(mlist.Machines, convertDBOtoMachineUIList(machine))
	}

	return mlist
}

type UIMachineList struct {
	Machines []UIMachine `json:"Machines"`
}

func convertDBOtoMachineUIList(inMachine Machine) UIMachine {
	var retMachine UIMachine

	retMachine.MachineID, _ = strconv.Atoi( inMachine.ID )
	retMachine.Token = inMachine.Token
	retMachine.Name = inMachine.MachineState.Alias
	retMachine.Firmware = inMachine.MachineState.CurrentFirmware
	retMachine.CurrentStatus = inMachine.CurrentStatus
	retMachine.SessionSinceClean = inMachine.MachineState.SessionStats.DirtySessionsSinceClean

	return retMachine
}

type UIMachine struct {
	MachineID         int    `json:"ID"`
	Name              string `json:"name"`
	Token             string `json:"token"`
	Firmware          string `json:"firmware"`
	SessionSinceClean int    `json:"usessinceclean"`
	CurrentStatus MachineStatus  `json:"CurrentStatus"`
}