package day14

import (
	"fmt"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// i assume this will be enough
const MAX_TREE_SECONDS = 10_000

func Part2(file *lib.InputFile) any {
	pRobots, height, width := parseInput(file)
	robots := *pRobots

	if len(robots) < 100 {
		// return "Example grid doesn't have an answer to part 2"
	}

	// 1:
	// it turns out when the tree is displayed
	// all the robots are on their own space
	// and have overlap otherwise
	// so you can just check for that

	// 2:
	// if that wasn't the case
	// you could check if there are enough robots near each other
	// see the commented out code below for that

	// 3:
	// once you know what you're looking for
	// you could also look for a long line of robots

	for second := 1; second <= MAX_TREE_SECONDS; second++ {

		// 1:
		visited := make([]bool, height * width)
		visitedCount := 0

		// 2:
		// visited := make([][]bool, height)
		// for i := range visited {
		// 	visited[i] = make([]bool, width)
		// }

		for i, robot := range robots {
			x := (robot.x + robot.vx) % width
			if x < 0 {
				x += width
			}

			y := (robot.y + robot.vy) % height
			if y < 0 {
				y += height
			}

			robots[i].x = x
			robots[i].y = y

			// 1:
			pos := y * width + x
			if !visited[pos] {
				visitedCount++
				visited[pos] = true
			}

			// 2:
			// if !visited[y][x] {
			// 	visited[y][x] = true
			// }
		}

		// 1:
		if visitedCount == len(robots) {
			return second
		}

		// 2:
		// robotsConnections := 0
		// for row, rowData := range visited {
		// 	for col, colData := range rowData {
		// 		if !colData {
		// 			continue
		// 		}

		// 		if row < len(visited)-1 && visited[row+1][col] {
		// 			robotsConnections++
		// 		}

		// 		if col < len(rowData)-1 && visited[row][col+1] {
		// 			robotsConnections++
		// 		}
		// 	}
		// }

		// if robotsConnections > len(robots) {
		// 	return second
		// }
	}

	return fmt.Sprintf("Didn't find any christmas tree within %d seconds", MAX_TREE_SECONDS)
}
