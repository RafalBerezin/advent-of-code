package day6

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	inputGrid := file.ByteGrid()

	height := len(inputGrid)
	width := len(inputGrid[0])

	guard := findGuard(inputGrid)

	// main logic reused in part 2
	// see shared.go
	visited := findVisitedCells(inputGrid, height, width, guard)

	result := 0
	for _, vis := range visited {
		if vis {
			result++
		}
	}

	return result
}
