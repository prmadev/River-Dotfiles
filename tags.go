package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func applyToTags(mwg *sync.WaitGroup) {
	for i := 1; i <= 9; i++ {
		numb := fmt.Sprint(i)
		tag := fmt.Sprint(1 << (i - 1))
		allCMDs := []*exec.Cmd{
			// focus on tag
			exec.Command(RIVERCTL, MAP, NORMAL, "Super", numb, "set-focused-tags", tag),
			exec.Command(RIVERCTL, MAP, NORMAL, "Super+Shift", numb, "set-view-tags", tag),
			// toggle focus of tag
			exec.Command(RIVERCTL, MAP, NORMAL, "Super+Control", numb, "toggle-focused-tags", tag),
			// toggle tag of view... I know!
			exec.Command(RIVERCTL, MAP, NORMAL, "Super+Shift+Control", numb, "toggle-view-tags", tag),
		}
		runner(allCMDs)

	}

	allTags := "$(((1 << 32) - 1))"
	allCMDs := []*exec.Cmd{
		// focus all tags
		exec.Command(RIVERCTL, MAP, NORMAL, "Super", "0", "set-focused-tags", allTags),
		// tag focused view with all tags (show on all tags)
		exec.Command(RIVERCTL, MAP, NORMAL, "Super+Shift", "0", "set-view-tags", allTags),
	}

	runner(allCMDs)

	mwg.Done()
}
