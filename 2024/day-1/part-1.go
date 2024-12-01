package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1() {
	ql := lib.NewQuickLogger(1, 1)
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

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	diffSum := 0
	for i := 0; i < length; i++ {
		diff := leftNums[i] - rightNums[i]
		if diff < 0 {
			diff = -diff
		}

		diffSum += diff
	}

	ql.Solve(diffSum)
}
