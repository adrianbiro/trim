package main

import "golang.org/x/term"

// todo flags for w and h
func ConsoleSize() (width int, height int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		width, height = 80, 20
	}
	// set half height or min
	switch {
	case height/2 > 20:
		height = height / 2
	default:
		height = 20
	}
	return
}
