package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
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
func killProcess(procName string) {

	p, err := exec.Command("pidof", "-S", "\n", procName).Output()
	if err != nil {
		return
	}
	// pp := strings.Split(string(p), "\n")
	pp := strings.Split(string(p), "\n")
	if len(pp[0]) == 0 {
		return
	}

	pid, err := strconv.Atoi(pp[0])
	if err != nil {
		log.Println("Err:", "trouble converting output of pidof:", err, "\n", string(p))
	}

	killErr := exec.Command("kill", fmt.Sprint(pid)).Run()
	if killErr != nil {
		log.Println("Err:", "having trouble killing", procName)
	}
}
