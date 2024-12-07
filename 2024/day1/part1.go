package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Strings()
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

	return diffSum
}
