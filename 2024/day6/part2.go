package day6

import (
	"sync"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// see shared.go

func Part2(file *lib.InputFile) any {
	inputGrid := file.ByteGrid()

	height := len(inputGrid)
	width := len(inputGrid[0])

	guard := findGuard(inputGrid)
	visited := findVisitedCells(inputGrid, height, width, guard)

	mut := sync.Mutex{}
	result := 0

	wg := sync.WaitGroup{}
	wg.Add(width * height)

	for row, rowSlice := range inputGrid  {
		for col := range rowSlice {
			go func() {
				if visited[row * width + col] && checkLoop(inputGrid, row, col, height, width, guard) {
					mut.Lock()
					result++
					mut.Unlock()
				}
				wg.Done()
			}()
		}
	}

	wg.Wait()

	return result
}

func checkLoop(grid [][]byte, row, col, height, width int, start lib.Point) bool {
	gridCopy := make([][]byte, height)
	for i, row := range grid {
		gridCopy[i] = make([]byte, width)
		copy(gridCopy[i], row)
	}

	dirI := 0
	dir := lib.Dirs4[dirI]
	guard := start

	gridCopy[row][col] = '#'
	visited := make([]bool, height * width * 4)

	for {
		unifiedPos := (guard.Y * width + guard.X) * 4 + dirI
		if visited[unifiedPos] {
			grid[row][col] = '.'
			return true
		}
		visited[unifiedPos] = true

		nextPos := guard.Add(&dir)
		if !lib.InBounds2D(nextPos.Y, nextPos.X, height, width) {
			return false
		}

		hitObstacle := gridCopy[nextPos.Y][nextPos.X] == '#'
		if hitObstacle {
			dirI = (dirI + 1) & 3
			dir = lib.Dirs4[dirI]
			continue
		}

		guard = nextPos
	}
}
