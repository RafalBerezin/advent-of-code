package day18

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)


func Part1(file *lib.InputFile) any {
	pPositions, size, steps := parseInput(file)
	positions := *pPositions

	grid := make([][]bool, size)
	for row := range grid {
		grid[row] = make([]bool, size)
	}

	for i := 0; i < steps; i++ {
		pos := positions[i]
		grid[pos.Y][pos.X] = true
	}

	result := findPath(&grid, size)

	if result == -1 {
		return "Could not find a valid path"
	}

	return result
}

// func NewGrid(parsed Parsed, length int) Grid {
// 	grid := Grid{Grid: make(map[Point]struct{})}
// 	for i, p := range parsed {
// 		if i == length {
// 			break
// 		}
// 		grid.Grid[p] = struct{}{}
// 		if p.X > grid.Max {
// 			grid.Max = p.X
// 		}
// 		if p.Y > grid.Max {
// 			grid.Max = p.Y
// 		}
// 	}
// 	return grid
// }
// func part1(parsed Parsed) {
// 	timeStart := time.Now()
// 	grid := NewGrid(parsed, 1024)
// 	grid.Print()
// 	steps := bfs(grid, Point{0, 0}, Point{grid.Max, grid.Max})
// 	fmt.Printf("Part 1: %d\t\tin %v\n", steps, time.Since(timeStart))
// }
// func part2(parsed Parsed) {
// 	timeStart := time.Now()
// 	gridFull := NewGrid(parsed, -1)
// 	grid := Grid{Grid: make(map[Point]struct{}), Max: gridFull.Max}
// 	for i, p := range parsed {
// 		grid.Grid[p] = struct{}{}
// 		steps := bfs(grid, Point{0, 0}, Point{gridFull.Max, gridFull.Max})
// 		if steps == -1 {
// 			grid.Print(p)
// 			fmt.Printf("Part 2: %d, @(%d,%d)\t\tin %v\n", i, p.X, p.Y, time.Since(timeStart))
// 			return
// 		}
// 	}
// 	fmt.Printf("Part 2: Not found\t\tin %v\n", time.Since(timeStart))
// }
