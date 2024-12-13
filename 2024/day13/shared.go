package day13

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

var ACost = 3
var BCost = 1

type machine struct {
	ax, ay, bx, by, prizeX, prizeY int
}

func parseInput(file *lib.InputFile) *[]machine {
	input := file.Strings()
	machines := make([]machine, (len(input) + 1) / 4)

	mi := 0
	for i := 0; i < len(input); i += 4 {
		aLine := input[i]
		ax, ay := parseLine(aLine, 2)

		bLine := input[i+1]
		bx, by := parseLine(bLine, 2)

		prizeLine := input[i+2]
		prizeX, prizeY := parseLine(prizeLine, 1)


		machines[mi] = machine{ax, ay, bx, by, prizeX, prizeY}
		mi++
	}

	return &machines
}


func parseLine(line string, skip int) (int, int) {
	parts := strings.Split(line, " ")[skip:]

	x, err := strconv.Atoi(parts[0][2 : len(parts[0]) - 1])
	lib.CheckError(err)

	y, err := strconv.Atoi(parts[1][2:])
	lib.CheckError(err)

	return x, y
}
