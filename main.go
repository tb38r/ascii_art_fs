package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1]

	if len(os.Args) != 3 || (os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy") {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	} else {

		file, err := os.Open(os.Args[2] + ".txt")
		if err != nil {
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("EX: go run . something standard")
			os.Exit(0)
		}
		defer file.Close()

		scanned := bufio.NewScanner(file) // reading file

		scanned.Split(bufio.ScanLines)

		var lines []string

		for scanned.Scan() {
			lines = append(lines, scanned.Text())
		}

		file.Close()

		asciiChrs := make(map[int][]string)
		id := 31

		for _, line := range lines {
			if string(line) == "" {
				id++
			} else {
				asciiChrs[id] = append(asciiChrs[id], line)
			}
		}

		for i := 0; i < len(args); i++ {
			if args[i] == 92 && args[i+1] == 110 {
				Newline(string(args[:i]), asciiChrs)
				Newline(string(args[i+2:]), asciiChrs)

			}
		}

		// checking for new line within arguments
		if !strings.Contains(args, "\\n") {
			Newline(args, asciiChrs)
		}

	}
}

func Newline(n string, y map[int][]string) {
	// prints horizontally
	for j := 0; j < len(y[32]); j++ {
		for _, letter := range n {
			fmt.Print(y[int(letter)][j])
		}
		fmt.Println()
	}
}
