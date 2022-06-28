package main

import "golang.org/x/term"

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

/* first draft with fallback executables
import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
// althw alternative executables and fallback values
func althw(cmdend string) int {
	cmd := exec.Command("tput", cmdend)
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	s := string(out)
	s = strings.TrimSpace(s)
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	if cmdend == "lines" {
		return 20
	}
	return 80

}

func oldconsoleSize() (int, int) {
	var (
		height, width int
	)
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil { //TODO
		log.Fatal(err)
	}
	s := string(out) // []uint8
	s = strings.TrimSpace(s)
	splits := strings.Split(s, " ")

	if height, err = strconv.Atoi(splits[0]); err != nil {
		//height = althw("lines")
		height = 20
	}
	if width, err = strconv.Atoi(splits[1]); err != nil {
		//width = althw("cols")
		width = 80
	}

	// half of terminal cap
	switch {
	case (height / 2) > 20:
		height = height / 2
	default:
		height = 20
	}
	return height, width
}
*/
