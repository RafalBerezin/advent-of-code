package day6

import (
	"slices"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

var chars = []byte{'^', '>', 'V', '<'}
var dirs = [][]int {
	{0,-1}, // ^
	{1,0}, // >
	{0,1}, // V
	{-1,0}, // <
}

func Part1(file *lib.InputFile) any {
	input := file.Bytes()

	width := slices.Index(input, '\n')
	widthNL := width + 1
	height := len(input) / widthNL

	startingPos := slices.IndexFunc(input, func(e byte) bool {
		return e == '^' || e == '>' || e == 'V' || e == '<'
	})

	x := startingPos % widthNL
	y := startingPos / widthNL

	charI := x + y * widthNL
	char := input[charI]
	dirI := slices.Index(chars, char)
	dir := dirs[dirI]

	for {
		input[charI] = '*'
		nextX := x + dir[0]
		nextY := y + dir[1]
		if nextX < 0 || nextX > width || nextY < 0 || nextY >= height {
			break
		}

		nextI := nextX + nextY * widthNL
		nextChar := input[nextI]

		if nextChar == '#' {
			dirI = (dirI + 1) % 4
			dir = dirs[dirI]
			continue
		}

		x = nextX
		y = nextY
		charI = nextI
	}

	result := 0
	for _, char := range input {
		if char == '*' {
			result++
		}
	}

	return result
}
