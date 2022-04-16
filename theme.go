package main

func SetTheme() {
	allArgs := [][]string{
		{"background-color", "0x191724"},
		{"border-color-focused", "0x31748f"},
		{"border-color-unfocused", "0x191724"},
		{"border-color-urgent", "0xeb6f92"},
		{"border-width", "10"},
		{"xcursor-theme", "'Layan-White Cursors'", "24"},
	}
	Riverctl(allArgs...)
}
