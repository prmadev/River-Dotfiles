package main

import (
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
	cmd.Run()
	wg.Done()
}

func cmdStart(cmd *exec.Cmd, swg *sync.WaitGroup) {
	cmd.Start()
	swg.Done()
}
