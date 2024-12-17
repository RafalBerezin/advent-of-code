package day17

import (
	"fmt"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	registers, instructions := parseInput(file)

	byteResults := runProgram(&registers, &instructions)

	stringResults := make([]string, len(byteResults))
	for i := range byteResults {
		stringResults[i] = fmt.Sprint(byteResults[i])
	}

	return strings.Join(stringResults, ",")
}
