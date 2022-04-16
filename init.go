package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	M  = "map"
	N  = "normal"
	MP = "map-pointer"
	SP = "spawn"
)

func main() {
	autorun()
	setTheme()
	setOptions()
	keyboardShortcuts()
	mouseShortCuts()
	assignToTags()
	inputs()
	finishingUp()
}
func inputs() {
	allArgs := [][]string{
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "events", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "drag", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "tap", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "natural-scroll", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "scroll-method", "two-finger"},
	}
	riverctl(allArgs...)
}
func setOptions() {
	allArgs := [][]string{
		// keyboard repeating rate
		{"set-repeat", "50", "300"},
		// Make certain views start floating
		{"float-filter-add", "app-id", "Rofi"},
		{"float-filter-add", "app-id", "float"},
		{"float-filter-add", "app-id", "popup"},
		{"float-filter-add", "app-id", "pinentry-qt"},
		{"float-filter-add", "app-id", "pinentry-gtk"},
		{"float-filter-add", "title", "Picture-in-Picture"},
		// Set app-ids and titles of views which should use client side decorations
		// {"csd-filter-add", "tapp-id", "\"gedit\""},
		// focus follows cursor
		{"focus-follows-cursor", "normal"},
		// cursor wrap
		{"set-cursor-warp", "on-output-change"},
		// attach mode
		{"attach-mode", "bottom"},
	}
	riverctl(allArgs...)
}

func setTheme() {
	allArgs := [][]string{
		{"background-color", "0x191724"},
		{"border-color-focused", "0x31748f"},
		{"border-color-unfocused", "0x191724"},
		{"border-color-urgent", "0xeb6f92"},
		{"border-width", "10"},
		{"xcursor-theme", "'Layan-White Cursors'", "24"},
	}
	riverctl(allArgs...)
}

func assignToTags() {
	for i := 1; i <= 9; i++ {
		numb := fmt.Sprint(i)
		tag := fmt.Sprint(1 << (i - 1))
		allArgs := [][]string{
			// focus on tag
			{M, N, "Super", numb, "set-focused-tags", tag},
			{M, N, "Super+Shift", numb, "set-view-tags", tag},
			// toggle focus of tag
			{M, N, "Super+Control", numb, "toggle-focused-tags", tag},
			// toggle tag of view... I know!
			{M, N, "Super+Shift+Control", numb, "toggle-view-tags", tag},
		}
		riverctl(allArgs...)
	}
	allTags := "$(((1 << 32) - 1))"
	allArgs := [][]string{
		// focus all tags
		{M, N, "Super", "0", "set-focused-tags", allTags},
		// tag focused view with all tags (show on all tags)
		{M, N, "Super+Shift", "0", "set-view-tags", allTags},
	}
	riverctl(allArgs...)
}

func mouseShortCuts() {
	allArgs := [][]string{
		{MP, N, "Super", "BTN_LEFT", "move-view"},
		{MP, N, "Super", "BTN_RIGHT", "resize-view"},
	}
	riverctl(allArgs...)
}

func keyboardShortcuts() {
	// default apps
	term := "kitty"
	browser := "qutebrowser"
	launcher := "rofi -show drun"
	// opening apps
	allArgs := [][]string{
		{M, N, "Super", "Return", SP, term},
		{M, N, "Super", "W", SP, browser},
		{M, N, "Super", "D", SP, launcher},
		// view focus control
		{M, N, "Super", "J", "focus-view", "next"},
		{M, N, "Super", "K", "focus-view", "previous"},
		// bump focused view to the top of the stack
		{M, N, "Super", "Space", "zoom"},
		{M, N, "Super", "Q", "close"},
		// output focus control
		{M, N, "Super", "Period", "focus-output", "next"},
		{M, N, "Super", "Comma", "focus-output", "previous"},
		// send view to output
		{M, N, "Super+Shift", "Period", "send-to-output", "next"},
		{M, N, "Super+Shift", "Comma", "send-to-output", "previous"},
		// resize the main ratio of rivertile
		{M, N, "Super", "H", "rivertile", "\"main-ratio -0.05\""},
		{M, N, "Super", "L", "rivertile", "\"main-ratio +0.05\""},
		// change the amount of views in main in rivertile
		{M, N, "Super+Shift", "H", "rivertile", "\"main-count +1\""},
		{M, N, "Super+Shift", "L", "rivertile", "\"main-count -1\""},
		// move views
		{M, N, "Super+Alt", "H", "move", "left", "100"},
		{M, N, "Super+Alt", "J", "move", "down", "100"},
		{M, N, "Super+Alt", "K", "move", "up", "100"},
		{M, N, "Super+Alt", "L", "move", "right", "100"},
		// snap views
		{M, N, "Super+Alt+Control", "H", "snap", "left"},
		{M, N, "Super+Alt+Control", "J", "snap", "down"},
		{M, N, "Super+Alt+Control", "K", "snap", "up"},
		{M, N, "Super+Alt+Control", "L", "snap", "right"},
		// move views
		{M, N, "Super+Alt+Shift", "H", "resize", "horizontal", "-100"},
		{M, N, "Super+Alt+Shift", "J", "resize", "vertical", "100"},
		{M, N, "Super+Alt+Shift", "K", "resize", "vertical", "-100"},
		{M, N, "Super+Alt+Shift", "L", "resize", "horizontal", "100"},
		// toggle layouts
		{M, N, "Super+Shift", "F", "toggle-float"},
		{M, N, "Super", "F", "toggle-fullscreen"},
		// Change layout orientation
		{M, N, "Super", "Up", "send-layout-cmd", "rivertile", "main-location", "top"},
		{M, N, "Super", "Right", "send-layout-cmd", "rivertile", "main-location", "right"},
		{M, N, "Super", "Down", "send-layout-cmd", "rivertile", "main-location", "bottom"},
		{M, N, "Super", "Left", "send-layout-cmd", "rivertile", "main-location", "left"},
		// media keys
		{M, N, "None", "XF86AudioMedia", SP, "\"playerctl play-pause\""},
		{M, N, "None", "XF86AudioPlay", SP, "\"playerctl play-pause\""},
		{M, N, "None", "XF86AudioPrev", SP, "\"playerctl previous\""},
		{M, N, "None", "XF86AudioNext", SP, "\"playerctl next\""},
		// volume keys
		{M, N, "None", "XF86AudioRaiseVolume", SP, "\"pactl set-sink-volume @DEFAULT_SINK@ +5%\""},
		{M, N, "None", "XF86AudioLowerVolume", SP, "\"pactl set-sink-volume @DEFAULT_SINK@ -5%\""},
		{M, N, "None", "XF86AudioMute", SP, "\"set-sink-mute @DEFAULT_SINK@ toggle\""},
		// brightness keys
		{M, N, "None", "XF86MonBrightnessUp", SP, "\"light -A 5\""},
		{M, N, "None", "XF86MonBrightnessDown", SP, "\"light -U 5\""},
	}
	riverctl(allArgs...)
}
func autorun() {
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
		exec.Command("wired","-r"),
	}
	for _, cmd := range cmdList {
		cmdStart(cmd)
	}
	waybarStart()
}

func waybarStart() {
	killCmd := exec.Command("killall", "waybar")
	cmdRun(killCmd)
	waybarCmd := exec.Command("waybar", "-c", "/home/amir/.config/river/waybar/config.json", "-s", "/home/amir/.config/river/waybar/style.css")
	cmdStart(waybarCmd)
}

func finishingUp() {
	cmd := exec.Command("rivertile", "-view-padding", "05", "-outer-padding", "05")
	cmdStart(cmd)
	allArgs := [][]string{
		{"default-layout", "rivertile"},
	}
	riverctl(allArgs...)
}

func riverctl(allArgs ...[]string) {
	for _, args := range allArgs {
		cmd := exec.Command("riverctl", args...)
		cmdRun(cmd)
	}
}

func cmdRun(cmd *exec.Cmd) {
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
