package day1

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
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

	rightCount := make(map[int]int)
	for _, num := range rightNums {
		rightCount[num]++
	}

	similarity := 0
	for _, num := range leftNums {
		similarity += num * rightCount[num]
	}

	return similarity
}
