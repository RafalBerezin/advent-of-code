package day10

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	grid := file.DigitGrid()

	height := len(grid)
	width := len(grid[0])

	result := 0
	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				result += calculateTrailRating(&grid, i, j, height, width)
			}
		}
	}

	return result
}

func calculateTrailRating(pGrid *[][]byte, row, col, height, width int) int {
	grid := *pGrid
	rating := 0

	var next func(current byte, y, x int)
	next = func(current byte, y, x int) {
		for _, dir := range lib.Dirs4 {
			newY := y + dir[0]
			newX := x + dir[1]

			if newY < 0 || newY >= height || newX < 0 || newX >= width {
				continue
			}

			nextCell := grid[newY][newX]
			if nextCell - current != 1 {
				continue
			}

			if nextCell == 9 {
				rating++
			} else {
				next(nextCell,  newY, newX)
			}
		}
	}

	next(0, row, col)

	return rating
}
