package main

func mouseBindings() {
	allArgs := [][]string{
		{MP, N, "Super", "BTN_LEFT", "move-view"},
		{MP, N, "Super", "BTN_RIGHT", "resize-view"},
	}
	riverctl(allArgs...)
}

func keyBindings() {
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
