package main

const (
	M  = "map"
	N  = "normal"
	MP = "map-pointer"
	SP = "spawn"
)

// main function Everything goes through here!
func main() {
	autorun()
	setTheme()
	setOptions()
	keyBindings()
	mouseBindings()
	makeTags()
	inputs()
	layout()
}
