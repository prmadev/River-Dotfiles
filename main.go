package main

import (
	"sync"
)

// main function Everything goes through here!
func main() {
	// for concurrency
	var mwg sync.WaitGroup

	mwg.Add(7) // if the number of gorouines under increased, this number should increase as well

	// some bindinggs
	go keyBindings(&mwg)
	go mouseBindings(&mwg)

	// coloring and stuff
	go setTheme(&mwg)

	// river's settings
	go setOptions(&mwg)
	go applyToTags(&mwg)
	go inputs(&mwg)

	//autorun
	go autorun(&mwg)

	// rest of concurrency stuff
	mwg.Wait()

}
