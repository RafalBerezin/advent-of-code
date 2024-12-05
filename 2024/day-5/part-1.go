package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1() {
	ql := lib.NewQuickLogger(5, 1)
	ql.Title()

	input, err := lib.LoadInputFile(5).Strings()
	lib.CheckError(err)
	splitIndex := slices.Index(input, "")

	orders := input[:splitIndex]
	prints := input[splitIndex+1:]

	ordersMap := make(map[string][]string)

	for _, order := range orders {
		o := strings.Split(order, "|")
		num := o[0]
		before := o[1]

		ordersMap[num] = append(ordersMap[num], before)
	}

	var result int
	for _, print := range prints {
		nums := strings.Split(print, ",")

		if isOrdered(nums, ordersMap) {
			middle := nums[len(nums)/2]

			num, err := strconv.Atoi(middle)
			lib.CheckError(err)

			result += num
		}
	}


	ql.Solve(result)
}

func isOrdered(nums []string, ordersMap map[string][]string) bool {
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		orderList := ordersMap[num]

		for j := 0; j < i; j++ {
			if slices.Contains(orderList, nums[j]) {
				return false
			}
		}
	}

	return true
}
