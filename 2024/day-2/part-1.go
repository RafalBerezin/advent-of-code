package main

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1() {
	ql := lib.NewQuickLogger(2, 1)
	ql.Title()

	input, err := lib.LoadInputFile(2).Strings()
	lib.CheckError(err)

	var safeReports int

	for _, report := range input {
		numStrs := strings.Split(report, " ")
		nums := make([]int, 0, len(numStrs))

		for _, str := range numStrs {
			num, err := strconv.Atoi(str)
			lib.CheckError(err)

			nums = append(nums, num)
		}

		firstDiff := nums[0] - nums[1]
		if firstDiff == 0 || firstDiff > 3 || firstDiff < -3 {
			continue
		}

		safe := true
		shouldAscend := firstDiff < 0

		maxLen := len(nums) - 1
		for i := 1; i < maxLen; i++ {
			diff := nums[i] - nums[i + 1]
			isAscending := diff < 0
			if diff == 0 || diff > 3 || diff < -3 || isAscending != shouldAscend {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}
	}

	ql.Solve(safeReports)
}
