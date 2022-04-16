package main

import "os/exec"

func Layout() {
	cmd := exec.Command("rivertile", "-view-padding", "05", "-outer-padding", "05")
	CmdStart(cmd)
	allArgs := [][]string{
		{"default-layout", "rivertile"},
	}
	Riverctl(allArgs...)
}
