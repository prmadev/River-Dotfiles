package main

import (
	"os/exec"
	"sync"
)

func inputs(mwg *sync.WaitGroup) {
	allCMDs := []*exec.Cmd{
		exec.Command(RIVERCTL, "input", "2:7:SynPS/2_Synaptics_TouchPad", "drag", "enabled"),
		exec.Command(RIVERCTL, "input", "2:7:SynPS/2_Synaptics_TouchPad", "tap", "enabled"),
		exec.Command(RIVERCTL, "input", "2:7:SynPS/2_Synaptics_TouchPad", "events", "enabled"),
		exec.Command(RIVERCTL, "input", "2:7:SynPS/2_Synaptics_TouchPad", "natural-scroll", "enabled"),
		exec.Command(RIVERCTL, "input", "2:7:SynPS/2_Synaptics_TouchPad", "scroll-method", "two-finger"),
	}
	runner(allCMDs)

	mwg.Done()
}
