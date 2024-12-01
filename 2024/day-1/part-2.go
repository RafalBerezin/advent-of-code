package main

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2() {
	ql := lib.NewQuickLogger(1, 2)
	ql.Title()

	input, err := lib.LoadInputFile(1).Strings()
	lib.CheckError(err)

	length := len(input)

	leftNums := make([]int, 0, length)
	rightNums := make([]int, 0, length)

	for _, str := range input {
		pair := strings.Split(str, "   ")

		left, err := strconv.Atoi(pair[0])
		lib.CheckError(err)
		right, err := strconv.Atoi(pair[1])
		lib.CheckError(err)

		leftNums = append(leftNums, left)
		rightNums = append(rightNums, right)
	}

	rightCount := make(map[int]int)
	for _, num := range rightNums {
		rightCount[num]++
	}

	similarity := 0
	for _, num := range leftNums {
		similarity += num * rightCount[num]
	}

	ql.Solve(similarity)
}
