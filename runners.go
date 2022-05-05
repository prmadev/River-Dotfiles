package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
)

func runner(allCMDs []*exec.Cmd) {
	var wg sync.WaitGroup

	for _, cmd := range allCMDs {
		wg.Add(1)
		go cmdRun(cmd, &wg)
	}

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

func cmdStart(cmd *exec.Cmd, swg *sync.WaitGroup) {
	var out bytes.Buffer

	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		swg.Done()
		return
	}
	swg.Done()
}
