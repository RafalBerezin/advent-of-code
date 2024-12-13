package day6

import "github.com/RafalBerezin/advent-of-code/2024/lib"

func findGuard(grid [][]byte) []int {
	for i, row := range grid {
		for j, cell := range row {
			if cell == '^' {
				return []int{i, j}
			}
		}
	}

	panic("Guard (^) not found.")
}

func findVisitedCells(grid [][]byte, height, width int, guard []int) []bool {
	dirI := 0 // up
	dir := lib.Dirs4[dirI]
	visited := make([]bool, height * width)

	for {
		visited[guard[0] * width + guard[1]] = true

		nextPos := []int{guard[0] + dir[0], guard[1] + dir[1]}
		if nextPos[0] < 0 || nextPos[0] >= height || nextPos[1] < 0 || nextPos[1] >= width {
			break
		}

		nextChar := grid[nextPos[0]][nextPos[1]]
		if nextChar == '#' {
			dirI = (dirI + 1) % 4
			dir = lib.Dirs4[dirI]
			continue
		}

		guard = nextPos
	}

	return visited
}
