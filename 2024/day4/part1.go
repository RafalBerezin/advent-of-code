package day4

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	result := 0
	for row, rowData := range grid {
		for col, char := range rowData {
			if char != 'X' {
				continue
			}

			for _, dir :=  range lib.Dirs8 {
				mRow, mCol := row + dir[0], col + dir[1]
				aRow, aCol := row + dir[0] * 2, col + dir[1] * 2
				sRow, sCol := row + dir[0] * 3, col + dir[1] * 3

				if !lib.InBounds2D(sRow, sCol, height, width) {
					continue
				}

				if grid[mRow][mCol] == 'M' && grid[aRow][aCol] == 'A' && grid[sRow][sCol] == 'S' {
					result++
				}
			}
		}
	}

	return result
}
