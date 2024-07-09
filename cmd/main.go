package main

import (
	. "henkel-somtec-devscript/env"
	. "henkel-somtec-devscript/persil"
	. "time"
)

func main() {
	CDCKey = "3_w-MmHQNyAXCXjY_ZuPuCY3qWdDw0qC0VnnnqXgt9hdgyjxAeyxoUVr822kQMjvZL"
	ApiURL = "https://app-api.stage.somtec.henkel.colorfy.cloud/v1/query"

	user := Login("", "")
	dev := user.ProvisionDevice()

	dev.Connect()

	dev.RemoveFromStation()

	dev.StartWashCycle()

	Sleep(3 * Second)

	dev.SetWashCycleStatus("washing")
	dev.SetBatteryLevel(80)
	dev.SetDetergentFillLevel(40)
	dev.SetWashLoad(20)

	Sleep(3 * Second)

	dev.SetWashCycleStatus("rinsing")

	Sleep(3 * Second)

	dev.StopWashCycle("done")

	dev.Disconnect()
}
