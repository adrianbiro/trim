package main

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

func consoleSize() (int, int) {
	var (
		heigth, width int
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

	if heigth, err = strconv.Atoi(splits[0]); err != nil {
		//heigth = althw("lines")
		heigth = 20
	}
	if width, err = strconv.Atoi(splits[1]); err != nil {
		//width = althw("cols")
		width = 80
	}

	// half of terminal cap
	switch {
	case (heigth / 2) > 20:
		heigth = heigth / 2
	default:
		heigth = 20
	}
	return heigth, width
}
