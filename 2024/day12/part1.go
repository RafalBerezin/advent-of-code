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

			p, a :=traverse(&grid, row, col, height, width, []int{0,0}, &areaMap)
			result += a*p
		}
	}

	return result
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
func traverse(pGrid *[][]byte, row, col, height, width int, from []int, pAreaMap *[]bool) (int, int) {
	grid := *pGrid
	areaMap := *pAreaMap
	privateAreaMap := make([]bool, height * width)

	currentType := grid[row][col]
	perimeter := 0

	var expand func(r, c  int)
	expand = func(r, c  int) {
		unifiedAreaPos := r * width + c
		if privateAreaMap[unifiedAreaPos] {
			return
		}
		privateAreaMap[unifiedAreaPos] = true

		for _, dir := range dirs {
			if dir[0] == from[0] && dir[1] == from[1] {
				continue
			}

			nextRow := r + dir[0]
			nextCol := c + dir[1]

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

	expand(row, col)

	area := 0
	for i, plot := range privateAreaMap {
		if plot {
			areaMap[i] = true
			area++
		}
	}

	return perimeter, area
}
