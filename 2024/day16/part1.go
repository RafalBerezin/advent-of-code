package day16

import (
	"math"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

const (
	startByte = 'S'
 	endByte = 'E'
	startingRotation = 1 // right

	moveCost = 1
	rotateCost = 1000
)

func Part1(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	startPos := lib.Point{X: -1, Y: -1}

	for row, rowData := range grid {
		for col, colData := range rowData {
			if colData == startByte {
				startPos.X = col
				startPos.Y = row
			}
		}
	}

	if startPos.X == -1 {
		return "Could not find starting positions"
	}

	result := findPath(&grid, height, width, startPos)

	return result
}

func findPath(pGrid *[][]byte, height, width int, start lib.Point) int {
	grid := *pGrid
	bestScore := math.MaxInt

	visited := make([][]int, height)
	for row := range visited {
		visited[row] = make([]int, width)
	}

	var traverse func(pos lib.Point, facing, score int)
	traverse = func(pos lib.Point, facing, score int) {
		dir := lib.Dirs4[facing]
		nextPos := pos.Add(&lib.Point{Y: dir[0], X: dir[1]})

		char := grid[nextPos.Y][nextPos.X]
		if char == '#' {
			return
		}

		posScore := visited[nextPos.Y][nextPos.X]
		if posScore != 0 && posScore < score {
			return
		}
		visited[nextPos.Y][nextPos.X] = score

		nextScore := score + moveCost
		if char == endByte {
			if nextScore < bestScore {
				bestScore = nextScore
			}
			return
		}

		if nextScore < bestScore {
			traverse(nextPos, facing, nextScore)
		}
		nextScore += rotateCost
		if nextScore < bestScore {
			traverse(nextPos, (facing - 1) & 3, nextScore)
			traverse(nextPos, (facing + 1) & 3, nextScore)
		}
	}

	traverse(start, startingRotation, 0)
	traverse(start, 0, 1000)

	return bestScore
}
