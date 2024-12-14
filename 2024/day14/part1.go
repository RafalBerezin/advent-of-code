package day14

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	robots, height, width := parseInput(file)

	centerHeight := height / 2
	centerWidth := width / 2

	quadrants := make([]int, 4)
	for _, robot := range *robots {
		x := (robot.x + robot.vx * 100) % width
		if x < 0 {
			x += width
		}

		y := (robot.y + robot.vy * 100) % height
		if y < 0 {
			y += height
		}

		dx := centerWidth - x
		dy := centerHeight - y

		if dx == 0 || dy == 0 {
			continue
		}

		quad := 0
		if dx < 0 {
			quad++
		}
		if dy < 0 {
			quad += 2
		}

		quadrants[quad]++
	}

	result := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]

	return result
}
