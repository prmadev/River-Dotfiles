package main

import "fmt"

func makeTags() {
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
