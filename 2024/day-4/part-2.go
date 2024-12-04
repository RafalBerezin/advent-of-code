package main

import (
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2() {
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
			if ch != "M" {
				continue
			}

			// forward
			if j + 2 < charsLen && i + 2 < linesLen {
				// right
				if checkMAS(input, i, j, false, true) {
					result++
				}
				// down
				if checkMAS(input, i, j, true, true) {
					result++
				}
			}
			// backwards
			if j >= 2 && i >= 2 {
				// left
				if checkMAS(input, i, j, false, false) {
					result++
				}
				// up
				if checkMAS(input, i, j, true, false) {
					result++
				}
			}
		}
	}

	ql.Solve(result)
}

func checkMAS(lines []string, row, col int, vertical, forward bool) bool {
	fD := -1
	if forward {
		fD = 1
	}
	fD2 := fD * 2

	vD := 2 * fD

	// middle A
	if charAt2(lines, row + fD, col + fD) != 'A' {
		return false
	}

	// other M
	if vertical {
		if charAt2(lines, row, col + vD) != 'M' {
			return false
		}
	} else {
		if charAt2(lines, row + vD, col) != 'M' {
			return false
		}
	}

	// diagonal S
	if charAt2(lines, row + fD*2, col + fD2) != 'S' {
		return false
	}

	// other S
	if vertical {
		return charAt2(lines, row + fD2, col) == 'S'
	} else {
		return charAt2(lines, row, col + fD2) == 'S'
	}
}

func charAt2(lines []string, row, col int) byte {
	return []byte(lines[row])[col]
}
