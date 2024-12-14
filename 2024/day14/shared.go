package day14

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

type robot struct {
	x, y, vx, vy int
}

const (
	WIDTH = 101
	HEIGHT = 103

	EXAMPLE_WIDTH = 11
	EXAMPLE_HEIGHT = 7
)


func parseInput(file *lib.InputFile) (*[]robot, int, int) {
	lines := file.Strings()
	robots := make([]robot, len(lines))

	height := HEIGHT
	width := WIDTH

	if len(lines) < 100 {
		height = EXAMPLE_HEIGHT
		width = EXAMPLE_WIDTH
	}

	for i, line := range lines {
		parts := strings.Split(line, " ")
		pos := strings.Split(parts[0][2:], ",")
		velocity := strings.Split(parts[1][2:], ",")

		x, err := strconv.Atoi(pos[0])
		lib.CheckError(err)
		y, err := strconv.Atoi(pos[1])
		lib.CheckError(err)

		vx, err := strconv.Atoi(velocity[0])
		lib.CheckError(err)
		vy, err := strconv.Atoi(velocity[1])
		lib.CheckError(err)

		robots[i] = robot{x, y, vx, vy}
	}

	return &robots, height, width
}
