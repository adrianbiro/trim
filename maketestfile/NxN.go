package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		cube, zfill    int
		arg, repstring string
	)
	switch {
	case len(os.Args[1:]) != 0:
		arg = os.Args[1]
		if num, err := strconv.Atoi(arg); err == nil {
			cube = num
			arg += " "
		}
	default:
		cube, arg = 200, "200 "
	}
	zfill = len(strconv.Itoa(cube))
	repstring = strings.Repeat(arg, cube)
	for i := 1; i <= cube; i++ {
		fmt.Printf("%0*d %v\n", zfill, i, repstring)
	}
}

/*
./NxN > text.txt 41K
./NxN 9876 > text.txt 466M
*/
