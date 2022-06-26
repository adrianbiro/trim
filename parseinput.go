package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

// towith return string adjusted to max width
func towidth(s string, w int) string {
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

// parseinput in height and width bounderies
func parseinput(f *os.File) {
	var (
		height, width          int    //= consoleSize()
		msg                    string = "more lines."
		counter                int
		linesleft, lenmsg      int
		repCount, lenlinesleft int
	)
	width, height, _ = term.GetSize(0) //todo err, fallback in file
	input := bufio.NewScanner(f)
	for input.Scan() {
		counter++
		if counter <= height {
			fmt.Println(towidth(input.Text(), width))
		}
	}
	// end block
	linesleft = counter - height
	lenlinesleft = len(string(linesleft))
	lenmsg = len(msg)
	repCount = (width - lenlinesleft - lenmsg - 2) // len of spaces
	if repCount < 0 {
		repCount = 0
	}
	fmt.Printf("%v%v %v\n", strings.Repeat(" ", repCount), linesleft, msg)
}

// FilesOrStdin read from files or stdin
func FilesOrStdin() {
	files := os.Args[1:]
	if len(files) == 0 {
		parseinput(os.Stdin) // FIXME e0xit status 1 file arg works
		// consoleSize() is the problem, init make no difference
		return
	}
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
			continue
		}
		parseinput(f)
		f.Close()
	}
}
