package day7

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Strings()
	result := 0

	for _, line := range input {
		separator := strings.Index(line, ":")
		if separator == -1 {
			continue
		}

		target, err := strconv.Atoi(line[:separator])
		lib.CheckError(err)

		numStrs := strings.Split(line[separator+2:], " ")
		nums := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			lib.CheckError(err)
			nums[i] = num
		}

		if checkNum(target, nums[0], nums, 1) {
			result += target
		}
	}

	return result
}

func checkNum(target, current int, nums []int, i int) bool {
	if i >= len(nums) {
		return current == target
	}
	next := nums[i]

	mul := current * next

	if mul <= target && checkNum(target, mul, nums, i + 1) {
		return true
	}

	add := current + next
	return add <= target && checkNum(target, add, nums, i + 1)
}

