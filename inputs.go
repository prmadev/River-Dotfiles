package main

func Inputs() {
	allArgs := [][]string{
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "events", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "drag", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "tap", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "natural-scroll", "enabled"},
		{"input", "2:7:SynPS/2_Synaptics_TouchPad", "scroll-method", "two-finger"},
	}
	Riverctl(allArgs...)
}
