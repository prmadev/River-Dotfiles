package main

import "os/exec"

func Autorun() {
	cmdList := []*exec.Cmd{
		exec.Command("swww", "init"),
		exec.Command("swww", "img", "/home/amir/.config/river/city.gif"),
		exec.Command("cfw"),
		exec.Command("udiskie"),
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
		exec.Command("wired", "-r"),
	}
	for _, cmd := range cmdList {
		CmdStart(cmd)
	}
	WaybarStart()
}

func WaybarStart() {
	killCmd := exec.Command("killall", "waybar")
	CmdRun(killCmd)
	waybarCmd := exec.Command("waybar", "-c", "/home/amir/.config/river/waybar/config.json", "-s", "/home/amir/.config/river/waybar/style.css")
	CmdStart(waybarCmd)
}
