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

			result += traverseSides(&grid, lib.Point{Y: row, X: col}, height, width, &visited)
		}
	}

	return result
}

func traverseSides(pGrid *[][]byte, start lib.Point, height, width int, pVisited *[]bool) int {
	grid := *pGrid
	visited := *pVisited

	areaType := grid[start.Y][start.X]
	area, sides := 0, 0

	var expand func(pos lib.Point)
	expand = func(pos lib.Point) {
		id := pos.Y * width + pos.X
		if visited[id] {
			return
		}
		visited[id] = true
		area++

		for i, dir := range lib.Dirs4 {
			nextPos := pos.Add(&dir)

			dirRotated := lib.Dirs4[(i + 1) & 3]
			nextPosRotated := pos.Add(&dirRotated)

			nextSameType :=
				lib.InBounds2D(nextPos.Y, nextPos.X, height, width) &&
				grid[nextPos.Y][nextPos.X] == areaType
			nextRotatedSameType :=
				lib.InBounds2D(nextPosRotated.Y, nextPosRotated.X, height, width) &&
				grid[nextPosRotated.Y][nextPosRotated.X] == areaType

			if !nextSameType && !nextRotatedSameType {
				sides++
			}
			nextPosDiagonal := nextPos.Add(&dirRotated)
			if nextSameType && nextRotatedSameType && grid[nextPosDiagonal.Y][nextPosDiagonal.X] != areaType {
				sides++
			}
		}

		for _, dir := range lib.Dirs4 {
			nextPos := pos.Add(&dir)

			if lib.InBounds2D(nextPos.Y, nextPos.X, height, width) && grid[nextPos.Y][nextPos.X] == areaType {
				expand(nextPos)
			}
		}
	}

	expand(start)

	return sides * area
}
