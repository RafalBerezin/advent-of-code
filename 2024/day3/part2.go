package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	input := file.Bytes()
	matcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do(n't)?\(\)`)

	enabled := true
	var result int

	matches := matcher.FindAll(input, -1)

	for _, match := range matches {
		if match[0] == byte('d') {
			enabled = len(match) == 4
			continue
		}

		if !enabled {
			continue
		}

		nums := strings.Split(string(match[4:len(match)-1]), ",")

		a, err := strconv.Atoi(nums[0])
		lib.CheckError(err)
		b, err := strconv.Atoi(nums[1])
		lib.CheckError(err)

		result += a * b
	}

	return result
}
