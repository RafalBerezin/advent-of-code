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

	robotRow := -1
	robotCol := -1
	for row, rowData := range input[:height] {
		warehouseGrid[row] = []byte(rowData)
		for col, colData := range warehouseGrid[row] {
			if colData == '@' {
				robotRow = row
				robotCol = col
			}
		}
	}

	if robotCol == -1 || robotRow == -1 {
		return "Could not find the robot (@)"
	}
	warehouseGrid[robotRow][robotCol] = '.'

	movements := []byte(strings.Join(input[height+1:], ""))

	for _, move := range movements {
		dir := lib.ByteDir(move)

		currentRow, currentCol := robotRow, robotCol
		boxes := 0
		for {
			nextRow := currentRow + dir[0]
			nextCol := currentCol + dir[1]
			nextSpace := warehouseGrid[nextRow][nextCol]

			if nextSpace == '#' {
				break
			}

			if nextSpace == 'O' {
				boxes++
				currentRow = nextRow
				currentCol = nextCol
				continue
			}

			if nextSpace == '.' {
				robotRow += dir[0]
				robotCol += dir[1]

				if boxes > 0 {
					warehouseGrid[robotRow][robotCol] = '.'
					warehouseGrid[nextRow][nextCol] = 'O'
				}

				break
			}
		}
	}

	warehouseGrid[robotRow][robotCol] = '@'

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
