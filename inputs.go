package main

import (
	"os/exec"
	"sync"
)

func inputs(mwg *sync.WaitGroup) {
	allCMDs := []*exec.Cmd{
		exec.Command(RIVERCTL, "input", "2:14:ETPS/2_Elantech_Touchpad", "drag", "enabled"),
		exec.Command(RIVERCTL, "input", "2:14:ETPS/2_Elantech_Touchpad", "tap", "enabled"),
		exec.Command(RIVERCTL, "input", "2:14:ETPS/2_Elantech_Touchpad", "events", "enabled"),
		exec.Command(RIVERCTL, "input", "2:14:ETPS/2_Elantech_Touchpad", "natural-scroll", "enabled"),
		exec.Command(RIVERCTL, "input", "2:14:ETPS/2_Elantech_Touchpad", "scroll-method", "two-finger"),
	}
	runner(allCMDs)

	mwg.Done()
}
