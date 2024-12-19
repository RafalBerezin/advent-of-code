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
				result += calculateTrailRating(&grid, lib.Point{Y: i, X: j}, height, width)
			}
		}
	}

	return result
}

func calculateTrailRating(pGrid *[][]byte, start lib.Point, height, width int) int {
	grid := *pGrid
	rating := 0

	var next func(current byte, pos lib.Point)
	next = func(current byte, pos lib.Point) {
		for _, dir := range lib.Dirs4 {
			newPos := pos.Add(&dir)

			if !lib.InBounds2D(newPos.Y, newPos.X, height, width) {
				continue
			}

			nextCell := grid[newPos.Y][newPos.X]
			if nextCell - current != 1 {
				continue
			}

			if nextCell == 9 {
				rating++
			} else {
				next(nextCell, newPos)
			}
		}
	}

	next(0, start)

	return rating
}
