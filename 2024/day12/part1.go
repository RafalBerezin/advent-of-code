package day12

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	result := 0
	visited := make([]bool, height * width)
	for row, rowData := range grid {
		for col := range rowData {
			if visited[row * width + col] {
				continue
			}

			result += traversePerimeter(&grid, row, col, height, width, &visited)
		}
	}

	return result
}

func traversePerimeter(pGrid *[][]byte, startRow, startCol, height, width int, pVisited *[]bool) int {
	grid := *pGrid
	visited := *pVisited

	areaType := grid[startRow][startCol]
	area, perimeter := 0, 0

	var expand func(row, col  int)
	expand = func(row, col  int) {
		pos := row * width + col
		if visited[pos] {
			return
		}
		visited[pos] = true
		area++

		for _, dir := range lib.Dirs4 {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			if !lib.InBounds2D(nextRow, nextCol, height, width) || grid[nextRow][nextCol] != areaType {
				perimeter++
				continue
			}

			expand(nextRow, nextCol)
		}
	}

	expand(startRow, startCol)

	return area * perimeter
}
