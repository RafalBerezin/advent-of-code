package day7

import (
	"strconv"
	"strings"
	"sync"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	input := file.Strings()

	mut := sync.Mutex{}
	result := 0

	wg := sync.WaitGroup{}
	wg.Add(len(input))

	for _, line := range input {
		go func() {
			defer wg.Done()

			separator := strings.Index(line, ":")
			if separator == -1 {
				return
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

			res := checkNum2(target, nums[0], nums, 1)
			if res {
				mut.Lock()
				result += target
				mut.Unlock()
			}
		}()
	}

	wg.Wait()

	return result
}

// done in 2 minutes then trying to find whats wrong
// inside i was calling chuckNum from part 1
func checkNum2(target, current int, nums []int, i int) bool {
	if i >= len(nums) {
		return current == target
	}
	next := nums[i]

	mul := current * next
	if mul <= target && checkNum2(target, mul, nums, i + 1) {
		return true
	}

	add := current + next
	if add <= target && checkNum2(target, add, nums, i + 1) {
		return true
	}

	concat, err := strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(next))
	lib.CheckError(err)

	return concat <= target && checkNum2(target, concat, nums, i + 1)
}

