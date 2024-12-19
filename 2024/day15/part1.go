package day15

import (
	"slices"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Strings()

	height := slices.Index(input, "")
	warehouseGrid := make([][]byte, height)

	robot := lib.Point{Y: -1, X: -1}
	for row, rowData := range input[:height] {
		warehouseGrid[row] = []byte(rowData)
		for col, colData := range warehouseGrid[row] {
			if colData == '@' {
				robot.Y = row
				robot.X = col
			}
		}
	}

	if robot.Y == -1 || robot.X == -1 {
		return "Could not find the robot (@)"
	}
	warehouseGrid[robot.Y][robot.X] = '.'

	movements := []byte(strings.Join(input[height+1:], ""))

	for _, move := range movements {
		dir := lib.ByteDir(move)

		pos := lib.Point{X: robot.X, Y: robot.Y}
		boxes := 0
		for {
			nextPos := pos.Add(&dir)
			nextSpace := warehouseGrid[nextPos.Y][nextPos.X]

			if nextSpace == '#' {
				break
			}

			if nextSpace == 'O' {
				boxes++
				pos = nextPos
				continue
			}

			if nextSpace == '.' {
				robot = robot.Add(&dir)

				if boxes > 0 {
					warehouseGrid[robot.Y][robot.X] = '.'
					warehouseGrid[nextPos.Y][nextPos.X] = 'O'
				}

				break
			}
		}
	}

	result := 0
	for row, rowData := range warehouseGrid {
		for col, colData := range rowData {
			if colData == 'O' {
				result += row * 100 + col
			}
		}
	}

	return result
}
