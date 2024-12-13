package day12

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
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

			result += traverseSides(&grid, row, col, height, width, &visited)
		}
	}

	return result
}

func traverseSides(pGrid *[][]byte, startRow, startCol, height, width int, pVisited *[]bool) int {
	grid := *pGrid
	visited := *pVisited

	areaType := grid[startRow][startCol]
	area, sides := 0, 0

	var expand func(row, col  int)
	expand = func(row, col  int) {
		pos := row * width + col
		if visited[pos] {
			return
		}
		visited[pos] = true
		area++

		for i, dir := range lib.Dirs4 {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			dir2 := lib.Dirs4[(i + 1) % 4]
			nextRowRotated := row + dir2[0]
			nextColRotated := col + dir2[1]

			nextSameType :=
				lib.InBounds2D(nextRow, nextCol, height, width) &&
				grid[nextRow][nextCol] == areaType
			nextRotatedSameType :=
				lib.InBounds2D(nextRowRotated, nextColRotated, height, width) &&
				grid[nextRowRotated][nextColRotated] == areaType

			if !nextSameType && !nextRotatedSameType {
				sides++
			}
			if nextSameType && nextRotatedSameType && grid[nextRow + dir2[0]][nextCol + dir2[1]] != areaType {
				sides++
			}
		}

		for _, dir := range lib.Dirs4 {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			if lib.InBounds2D(nextRow, nextCol, height, width) && grid[nextRow][nextCol] == areaType {
				expand(nextRow, nextCol)
			}
		}
	}

	expand(startRow, startCol)

	return sides * area
}
