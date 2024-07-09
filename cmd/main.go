package main

import (
	. "github.com/colorfy-software/henkel-somtec-devscript/env"
	. "github.com/colorfy-software/henkel-somtec-devscript/persil"
	. "time"
)

func main() {
	CDCKey = "3_w-MmHQNyAXCXjY_ZuPuCY3qWdDw0qC0VnnnqXgt9hdgyjxAeyxoUVr822kQMjvZL"
	ApiURL = "https://app-api.stage.somtec.henkel.colorfy.cloud/v1/query"

	user := Login("timo@colorfy.me", "]vW^a<f`7[d9V2en")
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
	dev.SetBatteryLevel(70)

	Sleep(3 * Second)

	dev.StopWashCycle("done")

	dev.ReturnToStation()

	dev.Disconnect()
}
