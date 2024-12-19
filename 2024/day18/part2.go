package day18

import (
	"fmt"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	pPositions, size, steps := parseInput(file)
	positions := *pPositions

	grid := make([][]bool, size)
	for row := range grid {
		grid[row] = make([]bool, size)
	}

	for i := 0; i < steps; i++ {
		pos := positions[i]
		grid[pos.Y][pos.X] = true
	}

	for i := steps; i < len(positions); i++ {
		nextPos := positions[i]
		grid[nextPos.Y][nextPos.X] = true

		gridCopy := *lib.CloneMatrix(&grid)
		if pathLen := findPath(&gridCopy, size); pathLen == -1 {
			return fmt.Sprintf("%d,%d", nextPos.X, nextPos.Y)
		}
	}

	return "Escape path never gets fully blocked"
}
