package main

import (
	"art"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1]

	if len(os.Args) != 3 || (os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy") {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Print("EX: go run . something standard")
		os.Exit(0)
	} else {

		file, err := os.Open(os.Args[2] + ".txt")
		if err != nil {
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Print("EX: go run . something standard")
			os.Exit(0)
		}
		defer file.Close()

		scanned := bufio.NewScanner(file) //reading file

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
				art.Newline(string(args[:i]), asciiChrs)
				art.Newline(string(args[i+2:]), asciiChrs)

			}
		}

		if strings.Contains(args, "\\n") == false { //checking for new line within arguments
			art.Newline(args, asciiChrs)
		}

	}
}
