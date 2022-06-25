package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		cube, zfill int
		arg         string
	)
	if len(os.Args[1:]) != 0 {
		arg = os.Args[1]
	}
	if num, err := strconv.Atoi(arg); err == nil {
		cube = num
	} else {
		cube = 200
	}
	zfill = len(strconv.Itoa(cube))
	arg = arg + " "
	for i := 1; i <= cube; i++ {
		fmt.Printf("%0*d %v\n", zfill, i, strings.Repeat(arg, cube))
	}
}

/*
./NxN.go > text.txt 41K
./NxN.go 9876 > text.txt 466M
*/
