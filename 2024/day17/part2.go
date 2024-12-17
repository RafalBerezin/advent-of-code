package day17

import (
	"slices"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// I assume that other input programs will have a similar structure
func Part2(file *lib.InputFile) any {
	_, instructions := parseInput(file)

	if res := tryNextBits(&instructions, 1, 0); res != -1 {
		return res
	}

	return "No valid solution found"
}

func tryNextBits(pInstructions *[]byte, insLen int, currentBits int) int {
	instructions := *pInstructions
	if insLen > len(instructions) {
		return currentBits
	}

	currentInstructions := instructions[len(instructions)-insLen:]

	for bits := 0; bits < 8; bits++ {
		nextBits := int((currentBits << 3) + int(bits))

		programResult := runProgram(&[]int{nextBits, 0, 0}, &instructions)
		if slices.Compare(programResult, currentInstructions) != 0 {
			continue
		}

		if result := tryNextBits(pInstructions, insLen+1, nextBits); result != -1 {
			return result
		}
	}

	return -1
}
