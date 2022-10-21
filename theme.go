package main

import (
	"os/exec"
	"sync"
)

func setTheme(mwg *sync.WaitGroup) {
	allCMDs := []*exec.Cmd{
		exec.Command(RIVERCTL, "background-color", "0x"+BASE),
		exec.Command(RIVERCTL, "border-color-focused", "0x"+LOVE),
		exec.Command(RIVERCTL, "border-color-unfocused", "0x"+BASE),
		exec.Command(RIVERCTL, "border-color-urgent", "0x"+LOVE),
		exec.Command(RIVERCTL, "border-width", "5"),
		exec.Command(RIVERCTL, "xcursor-theme", "'cutefish-light'", "24"),
	}

	runner(allCMDs)

	mwg.Done()
}

const (
	BASE    = "191724"
	SURFACE = "1F1D2E"
	OVERLAY = "26233A"
	MUTED   = "6E6A86"
	SUBTLE  = "908CAA"
	TEXT    = "E0DEF4"
	LOVE    = "EB6F92"
	GOLD    = "F6C177"
	ROSE    = "EBBCBA"
	PINE    = "31748F"
	FOAM    = "9CCFD8"
	IRIS    = "C4A7E7"
	HLLOW   = "21202E"
	HLMED   = "403D52"
	HLHIGH  = "524F67"
)
