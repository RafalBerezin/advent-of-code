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

			result += traversePerimeter(&grid, lib.Point{Y: row, X: col}, height, width, &visited)
		}
	}

	return result
}

func traversePerimeter(pGrid *[][]byte, start lib.Point, height, width int, pVisited *[]bool) int {
	grid := *pGrid
	visited := *pVisited

	areaType := grid[start.Y][start.X]
	area, perimeter := 0, 0

	var expand func(pos lib.Point)
	expand = func(pos lib.Point) {
		id := pos.Y * width + pos.X
		if visited[id] {
			return
		}
		visited[id] = true
		area++

		for _, dir := range lib.Dirs4 {
			nextPos := pos.Add(&dir)

			if !lib.InBounds2D(nextPos.Y, nextPos.X, height, width) || grid[nextPos.Y][nextPos.X] != areaType {
				perimeter++
				continue
			}

			expand(nextPos)
		}
	}

	expand(start)

	return area * perimeter
}
