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
	for i := 1; i <= cube; i++ {
		fmt.Println(fmt.Sprintf("%0*d", zfill, i), strings.Repeat("width ", cube))
	}
}

/*
go run NxN.go > text.txt
go run NxN.go num > text.txt
*/
