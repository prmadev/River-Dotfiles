package main

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
