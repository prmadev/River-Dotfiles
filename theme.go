package main

import (
	"os/exec"
	"sync"
)

func setTheme(mwg *sync.WaitGroup) {
	allCMDs := []*exec.Cmd{
		exec.Command(RIVERCTL, "background-color", "0x"+rosePine["bg"]),
		exec.Command(RIVERCTL, "border-color-focused", "0x"+rosePine["pine"]),
		exec.Command(RIVERCTL, "border-color-unfocused", "0x"+rosePine["bg"]),
		exec.Command(RIVERCTL, "border-color-urgent", "0x"+rosePine["love"]),
		exec.Command(RIVERCTL, "border-width", "6"),
		exec.Command(RIVERCTL, "xcursor-theme", "'cutefish-light'", "24"),
	}

	runner(allCMDs)

	mwg.Done()
}

var rosePine = map[string]string{
	"bg":    "E0E0E0",
	"surface": "E8E8E8",
	"overlay": "F0D0D0",
	"muted":   "585850",
	"subtle":  "404038",
	"text":    "505050",
	"love":    "806060",
	"gold":    "D8A838",
	"rose":    "F0D0D0",
	"pine":    "B0C8B0",
	"foam":    "9ccfd8",
	"iris":    "806060",
	"hlLow":   "D0D8D0",
	"hlMed":   "B0C0B0",
	"hlHigh":  "A0A098",
}

