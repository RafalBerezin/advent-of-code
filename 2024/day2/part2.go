package day2

import (
	"slices"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	input := file.Strings()

	var safeReports int

	for _, report := range input {
		numStrs := strings.Split(report, " ")
		nums := make([]int, 0, len(numStrs))

		for _, str := range numStrs {
			num, err := strconv.Atoi(str)
			lib.CheckError(err)

			nums = append(nums, num)
		}

		if checkNums(nums, true) {
			safeReports++
		}
	}

	return safeReports
}

func checkNums(nums []int, dampener bool) bool {
	// As it turns out I might be stupid
	// Lesson learned: Don't try to be smarter than you are
	// Make things work, then you can play around to make it better

	shouldAscend := nums[0] < nums[1]
	maxLen := len(nums) - 1

	for i := 0; i < maxLen; i++ {
		if isSafeStep(nums[i], nums[i+1], shouldAscend) {
			continue
		}

		if !dampener {
			return false
		}

		for j := 0; j < len(nums); j++ {
			if checkNums(slices.Concat(nums[:j], nums[j+1:]), false) {
				return true
			}
		}

		return false
	}

	return true
}

func isSafeStep(num1, num2 int, shouldAscend bool) bool {
	diff := num1 - num2
	isAscending := diff < 0
	return isAscending == shouldAscend && diff != 0 && diff <= 3 && diff >= -3
}
