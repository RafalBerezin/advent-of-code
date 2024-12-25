package day25

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

type pins [5]byte

func Part1(file *lib.InputFile) any {
	input := file.ByteGrid()

	locks := make([]pins, 0)
	keys := make([]pins, 0)

	for i := 0; i < len(input); i += 8 {
		inputType := input[i][0]

		var pins pins
		for pin := range pins {
			for row := 1; row <= 6; row ++ {
				if input[i+row][pin] != inputType {
					if inputType == '#' {
						pins[pin] = byte(row)-1
					} else {
						pins[pin] = 6 - byte(row)
					}
					break
				}
			}
		}

		if inputType == '#' {
			locks = append(locks, pins)
		} else {
			keys = append(keys, pins)
		}
	}

	result := 0
	for _, lock := range locks {
		middle: for _, key := range keys {
			for pin := range lock {
				if lock[pin] + key[pin] > 5 {
					continue middle
				}
			}
			result++
		}
	}

	return result
}

func Part2(file *lib.InputFile) any {
	return "Now all you need to do is to collect all previous 49 stars!"
}
