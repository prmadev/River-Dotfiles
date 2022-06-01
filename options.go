package main

import (
	"os/exec"
	"sync"
)

func setOptions(mwg *sync.WaitGroup) {
	allCMDs := []*exec.Cmd{
		// keyboard repeating rate
		exec.Command(RIVERCTL, "set-repeat", "50", "300"),

		// Make certain views start floating
		exec.Command(RIVERCTL, "float-filter-add", "app-id", "Rofi"),
		exec.Command(RIVERCTL, "float-filter-add", "app-id", "float"),
		exec.Command(RIVERCTL, "float-filter-add", "app-id", "popup"),
		exec.Command(RIVERCTL, "float-filter-add", "app-id", "pinentry-qt"),
		exec.Command(RIVERCTL, "float-filter-add", "app-id", "pinentry-gtk"),
		exec.Command(RIVERCTL, "float-filter-add", "title", "Picture-in-Picture"),
		exec.Command(RIVERCTL, "float-filter-add", "app-id", "launcher"),

		exec.Command(RIVERCTL, "csd-filter-add", "app-id", "Rofi"),
		exec.Command(RIVERCTL, "csd-filter-add", "app-id", "launcher"),
		// mouse stuff
		exec.Command(RIVERCTL, "focus-follows-cursor", "normal"),
		exec.Command(RIVERCTL, "set-cursor-warp", "on-output-change"),

		// layout related
		exec.Command(RIVERCTL, "attach-mode", "bottom"),       // new window's open at the end of stack instead of on top
		exec.Command(RIVERCTL, "default-layout", "rivertile"), // default layouting engine
	}

	runner(allCMDs)

	mwg.Done()
}
