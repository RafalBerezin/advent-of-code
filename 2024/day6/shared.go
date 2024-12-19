package day6

import "github.com/RafalBerezin/advent-of-code/2024/lib"

func findGuard(grid [][]byte) lib.Point {
	for row, rowData := range grid {
		for col, colData := range rowData {
			if colData == '^' {
				return lib.Point{X: col, Y: row}
			}
		}
	}

	panic("Guard (^) not found.")
}

func findVisitedCells(grid [][]byte, height, width int, guard lib.Point) []bool {
	dirI := 0 // up
	dir := lib.Dirs4[dirI]
	visited := make([]bool, height * width)

	for {
		visited[guard.Y * width + guard.X] = true

		nextPos := guard.Add(&dir)
		if !lib.InBounds2D(nextPos.Y, nextPos.X, height, width) {
			break
		}

		nextChar := grid[nextPos.Y][nextPos.X]
		if nextChar == '#' {
			dirI = (dirI + 1) & 3
			dir = lib.Dirs4[dirI]
			continue
		}

		guard = nextPos
	}

	return visited
}
