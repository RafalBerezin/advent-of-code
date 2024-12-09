package day9

import (
	"fmt"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Bytes()
	input = input[:len(input)-1] // remove new line

	fmt.Printf("spreadOut: %v\n", string(input))

	var spreadOut []int
	empty := false
	for i, char := range input {
		num := int(char - '0')
		toAppend := i/2 + '0'
		if empty {
			toAppend = -1
		}
		empty = !empty
		for j := 0; j < num; j++  {
			spreadOut = append(spreadOut, toAppend)
		}
	}

	fmt.Printf("spreadOut: %v\n", spreadOut)

	result := 0
	i := -1
	j := len(spreadOut)
	for {
		i++
		if i >= j {
			break
		}

		if spreadOut[i] != -1 {
			result += i * (int(spreadOut[i] - '0'))
			continue
		}

		for i < j {
			j--
			if spreadOut[j] != -1 {
				result += i * (int(spreadOut[j] - '0'))
				break
			}
		}
	}

	// result := Checksum("")

	return result
}
