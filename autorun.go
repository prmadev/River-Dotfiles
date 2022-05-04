package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
)

// autorun function ﳑ will start running everything at startup.
func autorun() {
	cmdList := []*exec.Cmd{
		// Setting wallpaper (use 'ln -P' to an image here.)
		exec.Command("swaybg", "-m", "fill", "-i", "/home/amir/.config/wallpaper"),

		// I really hate this thing. But for now I cannot replace it! :(
		// exec.Command("cfw"),
		exec.Command("tor"),
		// exec.Command("udiskie"),
		exec.Command(
			"ln",
			"-P",
			"--force",
			"/home/amir/.config/river/rofi/config.css",
			"/home/amir/.config/rofi/config.rasi",
		),
		exec.Command(
			"dbus-update-activation-environment",
			"SEATD_SOCK",
			"DISPLAY",
			"WAYLAND_DISPLAY",
			"XDG_SESSION_TYPE",
			"XDG_CURRENT_DESKTOP",
		),
		exec.Command("mako", "--default-timeout", "5000", "--background-color", "#31748f", "--border-color", "#e0def4", "--border-size", "2", "--font", "monospace", "--padding", "20", "--width", "350"),
	}

	for _, cmd := range cmdList {
		cmdStart(cmd)
	}

	waybarStart()
}

func waybarStart() {
	killCmd := exec.Command("killall", "waybar")

	var wg sync.WaitGroup
	wg.Add(1)
	go cmdRun(killCmd, &wg)
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

	waybarCmd := exec.Command("waybar", "-c", "/home/amir/.config/river/waybar/config.json", "-s", "/home/amir/.config/river/waybar/style.css")
	cmdStart(waybarCmd)
}
