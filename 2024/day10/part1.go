package day10

import "github.com/RafalBerezin/advent-of-code/2024/lib"

func Part1(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	trailHeads := make([][]bool, height)
	for i := range trailHeads {
		trailHeads[i] = make([]bool, width)
	}

	for i, row := range grid {
		for j, col := range row {
			if col == '0' {
				trailHeads[i][j] = true
			}
		}
	}

	result := 0
	for i, row := range trailHeads {
		for j, col := range row {
			if col {
				result += calculateTrailScore(&grid, i, j)
			}
		}
	}


	return result
}

func calculateTrailScore(pGrid *[][]byte, row, col int) int {
	grid := *pGrid
	height := len(grid)
	width := len(grid[0])

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	trailTails := make([]bool, height * width)
	var next func(current byte, y, x int)

	next = func(current byte, y, x int) {
		for _, dir := range dirs {
			newY := y + dir[0]
			newX := x + dir[1]

			if newY < 0 || newY >= height || newX < 0 || newX >= width {
				continue
			}

			nextCell := grid[newY][newX]
			if nextCell - current != 1 {
				continue
			}

			if nextCell == '9' {
				trailTails[newY * width + newX] = true
			} else {
				next(nextCell,  newY, newX)
			}
		}
	}

	next('0', row, col)

	score := 0
	for _, check := range trailTails {
		if check {
			score++
		}
	}

	return score
}
