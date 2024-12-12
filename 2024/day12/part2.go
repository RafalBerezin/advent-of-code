package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// really messy solution
// i'm going back to sleep
// i'll do something better when i wake up

func Part2(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	result := 0
	areaMap := make([]bool, height * width)
	for row, rowData := range grid {
		for col := range rowData {
			if areaMap[row * width + col] {
				continue
			}

			p, a :=traverse2(&grid, row, col, height, width, []int{0,0}, &areaMap)
			result += a*p
		}
	}

	return result
}

func traverse2(pGrid *[][]byte, row, col, height, width int, from []int, pAreaMap *[]bool) (int, int) {
	grid := *pGrid
	areaMap := *pAreaMap
	privateAreaMap := make([]bool, height * width)

	currentType := grid[row][col]
	sidesMap := make([]map[string]bool, 4)
	for i := range sidesMap {
		sidesMap[i] = make(map[string]bool)
	}

	var expand func(r, c  int)
	expand = func(r, c  int) {
		unifiedAreaPos := r * width + c
		if privateAreaMap[unifiedAreaPos] {
			return
		}
		privateAreaMap[unifiedAreaPos] = true

		for i, dir := range dirs {
			if dir[0] == from[0] && dir[1] == from[1] {
				continue
			}

			nextRow := r + dir[0]
			nextCol := c + dir[1]
			posStrKey := fmt.Sprintf("%d %d", nextRow, nextCol)

			if nextRow < 0 || nextRow >= height || nextCol < 0 || nextCol >= width {
				sidesMap[i][posStrKey] = true
				continue
			}

			nextType := grid[nextRow][nextCol]
			if nextType != currentType {
				sidesMap[i][posStrKey] = true
				continue
			}

			privateAreaMap[unifiedAreaPos] = true
			expand(nextRow, nextCol)
		}
	}

	expand(row, col)

	area := 0
	for i, plot := range privateAreaMap {
		if plot {
			areaMap[i] = true
			area++
		}
	}

	sides := 0
	for i, dir := range sidesMap {
		if len(dir) == 0 {
			continue
		}

		for {
			var element string
			for key := range dir {
				element = key
				break
			}
			if element == "" {
				break
			}
			delete(dir, element)

			parts := strings.Split(element, " ")
			row, err := strconv.Atoi(parts[0])
			lib.CheckError(err)
			col, err := strconv.Atoi(parts[1])
			lib.CheckError(err)

			sides++

			vertical := i & 1 == 1
			pRow, pCol := row, col

			for {
				pRow, pCol = newPos(pRow, pCol, true, vertical)
				pKey := fmt.Sprintf("%d %d", pRow, pCol)

				if _, found := dir[pKey]; found {
					delete(dir, pKey)
				} else {
					break
				}
			}

			nRow, nCol := row, col
			for {
				nRow, nCol = newPos(nRow, nCol, false, vertical)
				nKey := fmt.Sprintf("%d %d", nRow, nCol)

				if _, found := dir[nKey]; found {
					delete(dir, nKey)
				} else {
					break
				}
			}
		}
	}

	return sides, area
}

func newPos(row, col int, previous, vertical bool) (int, int) {
	d := 1
	if previous {
		d = -1
	}

	if vertical {
		return row + d, col
	}

	return row, col + d
}
