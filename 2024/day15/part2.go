package day15

import (
	"slices"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// i am not proud of what i've done
// but it works and i'm going back to sleep

func Part2(file *lib.InputFile) any {
	input := file.Strings()

	height := slices.Index(input, "")
	warehouseGrid := make([][]byte, height)

	robotRow := -1
	robotCol := -1
	for row, rowData := range input[:height] {
		warehouseGrid[row] = make([]byte, len(rowData) * 2)

		for col, colData := range []byte(rowData) {
			left := colData
			right := colData

			if colData == '@' {
				robotRow = row
				robotCol = col * 2
				right = '.'
			} else if colData == 'O' {
				left = '['
				right = ']'
			}

			warehouseGrid[row][col * 2] = left
			warehouseGrid[row][col * 2 + 1] = right
		}
	}

	if robotCol == -1 || robotRow == -1 {
		return "Could not find the robot (@)"
	}
	warehouseGrid[robotRow][robotCol] = '.'

	movements := []byte(strings.Join(input[height+1:], ""))

	for _, move := range movements {
		dir := lib.ByteDir(move)
		vertical := dir.Y != 0

		currentRow := robotRow
		boxes := 0

		if vertical {
			currentCols := [][]int{{robotCol}}
			i := 0
			outer: for {
				nextRow := currentRow + dir.Y
				nextCols := make([]int, 0)

				for _, currentCol := range currentCols[len(currentCols)-1] {
					nextSpace := warehouseGrid[nextRow][currentCol]

					if nextSpace == '#' {
						break outer
					}

					if nextSpace == '[' || nextSpace == ']' {
						otherCol := getOtherPosition(nextSpace, currentCol)

						if !slices.Contains(nextCols, currentCol) {
							nextCols = append(nextCols, currentCol)
						}
						if !slices.Contains(nextCols, otherCol) {
							nextCols = append(nextCols, otherCol)
						}
					}
				}

				if len(nextCols) == 0 {
					for i := len(currentCols) - 1; i >= 0; i-- {
						nextDI := (i+1) * dir.Y
						DI := i * dir.Y
						for _, col := range currentCols[i] {
							warehouseGrid[robotRow+nextDI][col] = warehouseGrid[robotRow+DI][col]
							warehouseGrid[robotRow+DI][col] = '.'
						}
					}
					robotRow += dir.Y
					break
				}

				currentRow = nextRow
				currentCols = append(currentCols, nextCols)
				i++
			}
		} else {
			currentCol := robotCol
			for {
				nextCol := currentCol + dir.X
				nextSpace := warehouseGrid[currentRow][nextCol]

				if nextSpace == '#' {
					break
				}

				if nextSpace == '[' || nextSpace == ']' {
					boxes += 2
					currentCol = nextCol + dir.X
					continue
				}

				if nextSpace == '.' {
					robotCol += dir.X

					if boxes > 0 {
						warehouseGrid[robotRow][robotCol] = '.'
						boxesCol := robotCol + dir.X
						boxes--

						from, to := boxesCol, boxesCol + boxes
						if dir.X == -1 {
							from, to = boxesCol - boxes, boxesCol
						}

						odd := from & 1
						for i := from; i <= to; i++ {
							if i&1 == odd {
								warehouseGrid[currentRow][i] = '['
							} else {
								warehouseGrid[currentRow][i] = ']'
							}
						}
					}

					break
				}
			}
		}
	}

	result := 0
	for row, rowData := range warehouseGrid {
		for col, colData := range rowData {
			if colData == '[' {
				result += row * 100 + col
			}
		}
	}

	return result
}

func getOtherPosition(boxPart byte, col int) int {
	if boxPart == '[' {
		return col + 1
	}
	return col - 1
}
