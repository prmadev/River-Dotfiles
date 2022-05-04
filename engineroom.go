package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"sync"
)

func riverctl(allArgs ...[]string) {
	var wg sync.WaitGroup
	for _, args := range allArgs {
		cmd := exec.Command("riverctl", args...)
		wg.Add(1)
		go cmdRun(cmd, &wg)
	}

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func cmdRun(cmd *exec.Cmd, wg *sync.WaitGroup) {
	var out bytes.Buffer

	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String() + cmd.String())
		wg.Done()
		return
	}

	wg.Done()
}

func cmdStart(cmd *exec.Cmd) {
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
