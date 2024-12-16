package day16

import (
	"math"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	startPos := lib.Point{X: -1, Y: -1}

	for row, rowData := range grid {
		for col, colData := range rowData {
			if colData == startByte {
				startPos.X = col-1
				startPos.Y = row
			}
		}
	}

	if startPos.X == -1 {
		return "Could not find starting positions"
	}

	result := findPathSeats(&grid, height, width, startPos)

	return result
}

func findPathSeats(pGrid *[][]byte, height, width int, start lib.Point) int {
	grid := *pGrid

	minScore := math.MaxInt
	var minVisits [][]bool
	visited := make([][]bool, height)
	for row := range visited {
		visited[row] = make([]bool, width)
	}

	posScores := make([][]int, height)
	for row := range posScores {
		posScores[row] = make([]int, width)
	}

	var traverse func(pos lib.Point, facing, score int)
	traverse = func(pos lib.Point, facing, score int) {
		dir := lib.Dirs4[facing]
		nextPos := pos.Add(&lib.Point{Y: dir[0], X: dir[1]})

		next := grid[nextPos.Y][nextPos.X]
		if next == '#' {
			return
		}

		posScore := posScores[nextPos.Y][nextPos.X]
		if posScore != 0 && posScore <= score {
			return
		}

		posScores[pos.Y][pos.X] = score
		visited[pos.Y][pos.X] = true
		defer func() {
			visited[pos.Y][pos.X] = false
		}()

		if next == endByte {
			if score < minScore {
				minScore = score
				minVisits = *lib.CloneMatrix(&visited)
				return
			}

			if score == minScore {
				for row, rowData := range visited {
					for col, colData := range rowData {
						if colData {
							minVisits[row][col] = true
						}
					}
				}
			}

			return
		}

		nextScore := score + moveCost
		if nextScore > minScore {
			return
		}
		traverse(nextPos, facing, score + 1);

		nextScore += rotateCost
		if nextScore <= minScore {
			traverse(nextPos, (facing + 1) & 3, score + 1001);
			traverse(nextPos, (facing - 1) & 3, score + 1001);
		}
	}

	traverse(start, startingRotation, 0)

	seats := 0
	for _, rowData := range minVisits {
		for _, colData := range rowData {
			if colData {
				seats++
			}
		}
	}

	return seats
}
