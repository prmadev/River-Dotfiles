package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Riverctl(allArgs ...[]string) {
	for _, args := range allArgs {
		cmd := exec.Command("riverctl", args...)
		CmdRun(cmd)
	}
}

func CmdRun(cmd *exec.Cmd) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String() + cmd.String())
		return
	}
}

func CmdStart(cmd *exec.Cmd) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}
