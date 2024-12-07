package main

import (
	"fmt"
	"slices"
	"sync"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

// chars and dirs in part-1.go

func Part2() {
	ql := lib.NewQuickLogger(6, 2)
	ql.Title()

	inputStrings, err := lib.LoadInputFile(6).Strings()
	lib.CheckError(err)

	height := len(inputStrings)
	width := len(inputStrings[0])

	var input [][]byte = make([][]byte, height)
	var guardPos []int = nil

	for row, line := range inputStrings {
		input[row] = []byte(line)

		if guardPos != nil {
			continue
		}

		col := slices.Index(input[row], '^')
		if col != -1 {
			guardPos = []int{col, row}
		}
	}


	result := 0
	wg := sync.WaitGroup{}
	wg.Add(width * height - 1)
	// no time to make it good, just check all options
	// i'll update this someday later
	for row := 0; row < height; row++  {
		for col := 0; col < width; col++  {
			go func() {
				if input[row][col] == '.' && 
				checkLoop(input, row, col, height, width, guardPos) {
					result++
				}
				wg.Done()
			}()
		}
	}

	wg.Wait()

	ql.Solve(result)
}

func checkLoop(input [][]byte, row, col, height, width int, start []int) bool {
	inputCopy := make([][]byte, height)
	for i, row := range input {
		inputCopy[i] = make([]byte, width)
		copy(inputCopy[i], row)
	}

	currentDirI := 0
	currentDir := dirs[currentDirI]
	currentPos := start

	inputCopy[row][col] = '#'

	visited := make(map[string]bool)

	for {
		visitedKey := fmt.Sprintf("%d,%d,%d", currentPos[0], currentPos[1], currentDirI)
		if visited[visitedKey] {
			return true
		}
		visited[visitedKey] = true

		nextPos := []int{
			currentPos[0] + currentDir[0],
			currentPos[1] + currentDir[1],
		}

		if nextPos[0] < 0 || nextPos[0] >= width || nextPos[1] < 0 || nextPos[1] >= height {
			return false
		}

		hitObstacle := inputCopy[nextPos[1]][nextPos[0]] == '#'
		if hitObstacle {
			currentDirI = (currentDirI + 1) % 4
			currentDir = dirs[currentDirI]
			continue
		}

		currentPos = nextPos
	}
}
