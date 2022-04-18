package main

import "os/exec"

func layout() {
	cmd := exec.Command("riverctl","spawn","'rivertile", "-view-padding", "05", "-outer-padding", "05'")
	cmdStart(cmd)
	allArgs := [][]string{
		{"default-layout", "rivertile"},
	}
	riverctl(allArgs...)
}
