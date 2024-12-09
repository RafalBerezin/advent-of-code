package day9

import (
	"fmt"
	"slices"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// Turns out reading the puzzle carefully might save you an hour
func Part2(file *lib.InputFile) any {
	input := file.Bytes()
	input = input[:len(input)-1] // remove new line

	fmt.Printf("spreadOut: %v\n", string(input))

	spreadOut := make([][]int, len(input))
	empty := true
	for i, char := range input {
		num := int(char - '0')
		empty = !empty

		if empty {
			spreadOut[i] = []int{num, -1}
			continue
		}

		spreadOut[i] = []int{num, i/2}
	}

	fmt.Printf("spreadOut: %v\n", spreadOut)

	for i := len(spreadOut) - 1; i > 0; i-- {
		current := spreadOut[i]
		cId := current[1]
		if cId == -1 {
			continue
		}

		cSize := current[0]
		for j := 0; j < i; j++ {
			slot := spreadOut[j]
			sId := slot[1]
			if sId != -1{
				continue
			}

			sSize := slot[0]
			if sSize < cSize {
				continue
			}

			fmt.Printf("Replacing %d: %v with %d %v\n", j, slot, i, current)

			sizeDiff := sSize - cSize
			slot[0] = current[0]
			slot[1] = current[1]
			current[1] = -1

			if sizeDiff > 0 {
				spreadOut = slices.Insert(spreadOut, j + 1, []int{sizeDiff, -1})
				i++
			}

			break
		}
	}

	result := 0
	totalI := 0
	for _, current := range spreadOut {
		size := current[0]
		id := current[1]

		if id == -1 {
			totalI += size
			continue
		}

		for k := 0; k < size; k++ {
			result += id * totalI
			totalI++
		}
	}

	fmt.Printf("spreadOut: %v\n", spreadOut)

	return result
}
