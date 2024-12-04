package main

import (
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1() {
	ql := lib.NewQuickLogger(4, 1)
	ql.Title()

	input, err := lib.LoadInputFile(4).Strings()
	lib.CheckError(err)

	linesLen := len(input)
	var result int
	for i, line := range input {
		chars := strings.Split(line, "")
		charsLen := len(chars)

		for j, ch := range chars {
			if ch != "X" {
				continue
			}

			// right
			if j + 3 < charsLen && checkDirection(input, i, j, 0, 1) {
				result++
			}
			// down
			if i + 3 < linesLen && checkDirection(input, i, j, 1, 0) {
				result++
			}
			// right down
			if j + 3 < charsLen && i + 3 < linesLen && checkDirection(input, i, j, 1, 1) {
				result++
			}
			// right up
			if j + 3 < charsLen && i >= 3 && checkDirection(input, i, j, -1, 1) {
				result++
			}
			// left
			if j >= 3 && checkDirection(input, i, j, 0, -1) {
				result++
			}
			// up
			if i >= 3 && checkDirection(input, i, j, -1, 0) {
				result++
			}
			// left up
			if j >= 3 && i >= 3 && checkDirection(input, i, j, -1, -1) {
				result++
			}
			// left down
			if j >= 3 && i + 3 < linesLen && checkDirection(input, i, j, 1, -1) {
				result++
			}
		}
	}

	ql.Solve(result)
}

func checkDirection(lines []string, row, col, rD, cD int) bool {
	return charAt(lines, row + rD, col + cD) == 'M' &&
		charAt(lines, row + rD*2, col + cD*2) == 'A' &&
		charAt(lines, row + rD*3, col + cD*3) == 'S'
}

func charAt(lines []string, row, col int) byte {
	return []byte(lines[row])[col]
}
