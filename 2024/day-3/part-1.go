package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1() {
	ql := lib.NewQuickLogger(3, 1)
	ql.Title()

	input, err := lib.LoadInputFile(3).Strings()
	lib.CheckError(err)

	matcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	var result int
	for _, line := range input {
		matches := matcher.FindAll([]byte(line), -1)

		for _, match := range matches {
			nums := strings.Split(string(match[4:len(match)-1]), ",")

			a, err := strconv.Atoi(nums[0])
			lib.CheckError(err)
			b, err := strconv.Atoi(nums[1])
			lib.CheckError(err)

			result += a * b
		}
	}

	ql.Solve(result)
}
