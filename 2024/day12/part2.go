package day12

import (
	"sort"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

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

			sides, area := traverseSides(&grid, row, col, height, width, &areaMap)
			result += sides * area
		}
	}

	return result
}

func traverseSides(pGrid *[][]byte, startRow, startCol, height, width int, pAreaMap *[]bool) (int, int) {
	grid := *pGrid
	areaMap := *pAreaMap

	areaType := grid[startRow][startCol]
	localAreaMap := make([]bool, height * width)

	areaSidesMap := make([][][]int, 4)
	for i := range areaSidesMap {
		length := height
		if i & 1 == 1 {
			length = width
		}
		areaSidesMap[i] = make([][]int, length + 2)
	}

	var expand func(row, col  int)
	expand = func(row, col  int) {
		unifiedAreaPos := row * width + col
		if localAreaMap[unifiedAreaPos] {
			return
		}
		localAreaMap[unifiedAreaPos] = true

		for i, dir := range dirs {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			if nextRow < 0 || nextRow >= height || nextCol < 0 || nextCol >= width {
				storeSide(&areaSidesMap, row, col, i)
				continue
			}

			nextType := grid[nextRow][nextCol]
			if nextType != areaType {
				storeSide(&areaSidesMap, row, col, i)
				continue
			}

			localAreaMap[unifiedAreaPos] = true
			expand(nextRow, nextCol)
		}
	}

	expand(startRow, startCol)

	area := 0
	for i, plot := range localAreaMap {
		if plot {
			areaMap[i] = true
			area++
		}
	}

	sides := 0
	for _, dir := range areaSidesMap {
		for _, edgeGroup := range dir {
			if len(edgeGroup) == 0 {
				continue
			}

			sort.Ints(edgeGroup)

			edge := -10 // the lowest number that should appear is -1
			for i := 0; i < len(edgeGroup); i++ {
				nextEdge := edgeGroup[i]
				if edge + 1 != nextEdge {
					sides++
				}
				edge = nextEdge
			}
		}
	}

	return sides, area
}

func storeSide(pSidesMap *[][][]int, row, col, dirI int) {
	sidesMap := *pSidesMap

	verticalEdge := dirI & 1 == 1
	edgeGroup := row
	edgePlace := col
	if verticalEdge {
		edgeGroup = col
		edgePlace = row
	}

	sidesMap[dirI][edgeGroup] = append(sidesMap[dirI][edgeGroup], edgePlace)
}
