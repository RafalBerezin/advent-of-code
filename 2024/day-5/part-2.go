package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2() {
	ql := lib.NewQuickLogger(5, 2)
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

		// isOrdered func from part-1.go
		if isOrdered(nums, ordersMap) {
			continue
		}

		slices.SortFunc(nums, func(a, b string) int {
			if slices.Contains(ordersMap[a], b) {
				return 1
			} else {
				return -1
			}
		})

		middle := nums[len(nums)/2]

		num, err := strconv.Atoi(middle)
		lib.CheckError(err)

		result += num
	}


	ql.Solve(result)
}
