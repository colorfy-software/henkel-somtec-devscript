package main

import (
	. "github.com/colorfy-software/henkel-somtec-devscript/persil"
	. "time"
)

func main() {
	// Set Environment to pilot
	SetEnv(Pilot)

	// Login with user and provision new device
	user := Login("timo@colorfy.me", "]vW^a<f`7[d9V2en")
	dev := user.ProvisionDevice()

	// Connect device to backend
	dev.Connect()

	// Remove device from station by setting its charging status to 'discharging'
	dev.SetChargingStatus("discharging")

	// Start wash cycle
	dev.StartWashCycle()

	// Wait for 3 seconds
	Sleep(3 * Second)

	// Set some device properties
	dev.SetWashCycleStatus("washing")
	dev.SetBatteryLevel(80)
	dev.SetDetergentFillLevel(40)
	dev.SetWashLoad(20)

	// Wait for 3 seconds
	Sleep(3 * Second)

	// Set some more device properties
	dev.SetWashCycleStatus("rinsing")
	dev.SetBatteryLevel(70)

	// Wait for 3 seconds
	Sleep(3 * Second)

	// Stop wash cycle with status 'done'
	dev.StopWashCycle("done")

	// Put device back on station
	dev.SetChargingStatus("charging")

	// Disconnect device from backend
	dev.Disconnect()
}
