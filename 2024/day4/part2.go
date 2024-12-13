package day4

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	result := 0
	for row := 1; row < height - 1; row++ {
		for col := 1; col < width - 1; col++ {
			if grid[row][col] != 'A' {
				continue
			}

			chars := make([]byte, 4)
			for i, dir :=  range lib.Dirs4Diagonal {
				chars[i] = grid[row + dir[0]][col + dir[1]]
			}

			for i := range chars {
				if chars[i] == 'M' &&
					chars[(i+1)&3] == 'M' &&
					chars[(i+2)&3] == 'S' &&
					chars[(i+3)&3] == 'S' {
						result++
						break
					}
			}
		}
	}

	return result
}
