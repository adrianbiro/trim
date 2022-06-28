package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// towith return string adjusted to max width
func ToWidth(s string, w int) string {
	if len(s) > w {
		w -= 3
		var builder strings.Builder
		for i, j := range s {
			switch {
			case i+1 < w:
				builder.WriteString(string(j))
			default:
				builder.WriteString("...")
				return builder.String()
			}
		}
	}
	return s
}

// ParseInput in height and width bounderies
func ParseInput(f *os.File) {
	var (
		width, height int = ConsoleSize()
		counter       int
	)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counter++
		if counter <= height {
			fmt.Println(ToWidth(input.Text(), width))
		}
	}
	// end block
	if counter > height {
		msg := "more lines."
		lenmsg := len(msg)
		linesleft := counter - height
		lenlinesleft := len(string(rune(linesleft)))
		repCount := (width - lenlinesleft - lenmsg - 2) // len of spaces
		if repCount < 0 {
			repCount = 0
		}
		fmt.Printf("%v%v %v\n", strings.Repeat(" ", repCount), linesleft, msg)
	}
	fmt.Printf("%v lines in total.\n", counter)
}

// FilesOrStdin read from files or stdin
func FilesOrStdin() {
	files := os.Args[1:]
	if len(files) == 0 {
		ParseInput(os.Stdin) // TODO term.GetSize() return -1 -1 err when stdin
		return
	}
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
			continue
		}
		ParseInput(f)
		f.Close()
	}
}
