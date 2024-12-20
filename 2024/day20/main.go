package day20

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

const (
	MAX_CHEAT_1 = 2
	MAX_CHEAT_2 = 20

	MIN_SAVE = 100
	EXAMPLE_MIN_SAVE = 50
)

func Part1(file *lib.InputFile) any {
	stepScores, minSave := parseInput(file)
	return findCheats(stepScores, MAX_CHEAT_1, minSave)
}

func Part2(file *lib.InputFile) any {
	stepScores, minSave := parseInput(file)
	return findCheats(stepScores, MAX_CHEAT_2, minSave)
}

func parseInput(file *lib.InputFile) (*[][]int, int) {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	start, end := lib.Point{}, lib.Point{}
	for row, rowData := range grid {
		for col, colData := range rowData {
			if colData == 'S' {
				start = lib.Point{Y: row, X: col}
				continue
			}

			if colData == 'E' {
				end = lib.Point{Y: row, X: col}
				continue
			}
		}
	}

	// the steps are 1 higher than in reality
	// to mark the starting position as non zero
	// because the cheat search treats 0's as not on the path
	// we only care about those relative to each other anyway
	stepScores := make([][]int, height)
	for row := range stepScores {
		stepScores[row] = make([]int, width)
	}

	stepScores[start.Y][start.X] = 1
	steps := 1

	current := start
	previous := start

	for current != end {
		steps++

		for _, dir := range lib.Dirs4 {
			next := current.Add(&dir)
			if next == previous {
				continue
			}
			if grid[next.Y][next.X] != '#' {
				stepScores[next.Y][next.X] = steps

				previous = current
				current = next

				break
			}
		}
	}

	minSave := MIN_SAVE
	if height < 50 {
		minSave = EXAMPLE_MIN_SAVE
	}

	return &stepScores, minSave
}

func findCheats(pStepScores *[][]int, maxCheatDistance, minStepsSave int) int {
	stepScores := *pStepScores

	height := len(stepScores)
	width := len(stepScores[0])

	viableCheats := 0

	for row, rowData := range stepScores {
		for col, currentSteps := range rowData {
			if currentSteps == 0 {
				continue
			}

			for dRow := -maxCheatDistance; dRow <= maxCheatDistance; dRow++ {
				absDRow := dRow
				if absDRow < 0 {
					absDRow = -absDRow
				}

				maxDCol := maxCheatDistance - absDRow
				for dCol := -maxDCol; dCol <= maxDCol; dCol++ {
					absDCol := dCol
					if absDCol < 0 {
						absDCol= -absDCol
					}

					targetRow := row + dRow
					targetCol := col + dCol

					if !lib.InBounds2D(targetRow, targetCol, height, width) {
						continue
					}

					targetSteps := stepScores[targetRow][targetCol]
					if targetSteps - currentSteps - absDRow - absDCol >= minStepsSave {
						viableCheats++
					}
				}
			}
		}
	}

	return viableCheats
}
