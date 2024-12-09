package day9

import (
	"slices"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// Turns out reading the puzzle carefully might save you an hour
func Part2(file *lib.InputFile) any {
	input := file.Digits()

	unfoldedMap := make([][]int, len(input))
	for i, size := range input {
		if i&1 == 1 {
			unfoldedMap[i] = []int{size, -1}
			continue
		}

		unfoldedMap[i] = []int{size, i/2}
	}

	for i := len(unfoldedMap) - 1; i > 0; i-- {
		current := unfoldedMap[i]
		cId := current[1]
		if cId == -1 {
			continue
		}

		cSize := current[0]
		for j := 0; j < i; j++ {
			slot := unfoldedMap[j]
			sId := slot[1]
			if sId != -1{
				continue
			}

			sSize := slot[0]
			if sSize < cSize {
				continue
			}

			sizeDiff := sSize - cSize
			slot[0] = current[0]
			slot[1] = current[1]
			current[1] = -1

			if sizeDiff > 0 {
				unfoldedMap = slices.Insert(unfoldedMap, j + 1, []int{sizeDiff, -1})
				i++
			}

			break
		}
	}

	result := 0
	totalI := 0
	for _, current := range unfoldedMap {
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

	return result
}
