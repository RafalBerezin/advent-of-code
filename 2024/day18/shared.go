package day18

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

const (
	EXAMPLE_GRID_SIZE = 7
	GRID_SIZE = 71

	EXAMPLE_STEPS = 12
	STEPS = 1024
)

func parseInput(file *lib.InputFile) (*[]lib.Point, int, int) {
	input := file.Strings()

	size := GRID_SIZE
	steps := STEPS
	if len(input) < 100 {
		size = EXAMPLE_GRID_SIZE
		steps = EXAMPLE_STEPS
	}

	positions := make([]lib.Point, len(input))
	for i, line := range input {
		parts := strings.Split(line, ",")

		x, err := strconv.Atoi(parts[0])
		lib.CheckError(err)
		y, err := strconv.Atoi(parts[1])
		lib.CheckError(err)

		positions[i] = lib.Point{X: x, Y: y}
	}

	return &positions, size, steps
}

func findPath(pGrid *[][]bool, size int) int {
	visited := *pGrid

	end := lib.Point{Y: size - 1, X: size - 1}

	steps := 0
	queue := []lib.Point{{Y: 0, X: 0}}

	for len(queue) > 0 {
		nextQueue := make([]lib.Point, 0)

		for _, pos := range queue {
			if pos == end {
				return steps
			}

			if visited[pos.Y][pos.X] {
				continue
			}
			visited[pos.Y][pos.X] = true

			for _, dir := range lib.Dirs4 {
				nextPos := pos.Add(&dir)

				if lib.InBounds2D(nextPos.Y, nextPos.X, size, size) {
					nextQueue = append(nextQueue, nextPos)
				}
			}
		}

		queue = nextQueue
		steps ++
	}

	return -1
}
