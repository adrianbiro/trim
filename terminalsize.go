package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func consoleSize() (int, int) {
	cmd := exec.Command("stty", "size")
	// fall back cmd and values 10 80
	//cmd := exec.Command("stput" "cols")
	//cmd := exec.Command("stput" "lines")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err) // TODO  Errorf(format string, a ...any) error or default values
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return heigth, width
}
