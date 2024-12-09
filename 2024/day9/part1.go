package day9

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Digits()

	diskSize := 0
	for _, digit := range input {
		diskSize += digit
	}

	unfoldedMap := make([]int, diskSize)
	resultI := 0
	for i, size := range input {
		id := -1
		if i&1 == 0 {
			id = i/2
		}

		for j := 0; j < size; j++ {
			unfoldedMap[resultI + j] = id
		}
		resultI += size
	}

	result := 0
	low, high := -1, diskSize
	for {
		low++
		if low >= high {
			break
		}

		if unfoldedMap[low] != -1 {
			result += low * unfoldedMap[low]
			continue
		}

		for low < high {
			high--
			if unfoldedMap[high] != -1 {
				result += low * unfoldedMap[high]
				break
			}
		}
	}

	return result
}
