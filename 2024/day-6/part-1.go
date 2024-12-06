package main

import (
	"fmt"
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

func Part1() {
	ql := lib.NewQuickLogger(6, 1)
	ql.Title()

	input := lib.LoadInputFile(6).Bytes()

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
		fmt.Printf("dir: %v\n", dir)
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

	fmt.Printf("path:\n%s\n", input)

	result := 0
	for _, char := range input {
		if char == '*' {
			result++
		}
	}

	ql.Solve(result)
}
