package day6

var dirs = [][]int {
	{-1,0}, // up
	{0,1}, // right
	{1,0}, // down
	{0,-1}, // left
}

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
	dir := dirs[dirI]
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
			dir = dirs[dirI]
			continue
		}

		guard = nextPos
	}

	return visited
}
