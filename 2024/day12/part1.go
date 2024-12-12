package day12

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	result := 0
	areaMap := make([]bool, height * width)
	for row, rowData := range grid {
		for col := range rowData {
			if areaMap[row * width + col] {
				continue
			}

			perimeter, area := traversePerimeter(&grid, row, col, height, width, &areaMap)
			result += perimeter * area
		}
	}

	return result
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
func traversePerimeter(pGrid *[][]byte, startRow, startCol, height, width int, pAreaMap *[]bool) (int, int) {
	grid := *pGrid
	areaMap := *pAreaMap
	privateAreaMap := make([]bool, height * width)

	currentType := grid[startRow][startCol]
	perimeter := 0

	var expand func(row, col  int)
	expand = func(row, col  int) {
		unifiedAreaPos := row * width + col
		if privateAreaMap[unifiedAreaPos] {
			return
		}
		privateAreaMap[unifiedAreaPos] = true

		for _, dir := range dirs {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			if nextRow < 0 || nextRow >= height || nextCol < 0 || nextCol >= width {
				perimeter++
				continue
			}

			nextType := grid[nextRow][nextCol]
			if nextType != currentType {
				perimeter++
				continue
			}

			privateAreaMap[unifiedAreaPos] = true
			expand(nextRow, nextCol)
		}
	}

	expand(startRow, startCol)

	area := 0
	for i, plot := range privateAreaMap {
		if plot {
			areaMap[i] = true
			area++
		}
	}

	return perimeter, area
}
