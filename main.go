package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// towith return string adjusted to max width TODO unicode works?
func towidth(s string, w int) string {
	if len(s) > w {
		w -= 3
		var builder strings.Builder
		builder.WriteString(s[:w])
		builder.WriteString("...")
		return builder.String()
	}
	return s
}

// parseinput in height and width bounderies
func parseinput(f *os.File) {
	var (
		height, width int    = consoleSize()
		msg           string = "more lines."
		counter       int
		linesleft     int
		repCount      int
	)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counter++
		if counter <= height {
			fmt.Println(towidth(input.Text(), width))
		}

	}
	// end block
	linesleft = counter - height
	repCount = (width - len(string(linesleft)) - len(msg) - 2)
	if repCount < 0 {
		repCount = 0
	}
	fmt.Printf("%v%v %v\n", strings.Repeat(" ", repCount), linesleft, msg)
}

// FilesOrStdin read from files or stdin
func FilesOrStdin() {
	files := os.Args[1:]
	if len(files) == 0 {
		parseinput(os.Stdin) // FIXME exit status 1
	} else {
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
}

func main() {
	FilesOrStdin()
}
