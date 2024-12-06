package main

import (
	"fmt"
	"os"
)

func main() {
	argsCount := len(os.Args)
	if (argsCount < 2) {
		Part1()
		Part2()
		return
	}

	part := os.Args[1]
	if part != "1" && part != "2" {
		fmt.Println("specified part must be either 1 or 2")
		return
	}

	if part == "1" {
		Part1()
	} else {
		Part2()
	}
}
