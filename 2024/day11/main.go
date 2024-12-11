package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	return iterateAndCount(file, 25)
}

func Part2(file *lib.InputFile) any {
	return iterateAndCount(file, 75)
}


func iterateAndCount(file *lib.InputFile, iterations int) int {
	input := file.Strings()
	stoneStrs := strings.Split(input[0], " ")

	stoneCounts := make(map[int]int)
	for _, stoneStr := range stoneStrs {
		stone, err := strconv.Atoi(stoneStr)
		lib.CheckError(err)
		stoneCounts[stone]++
	}

	for blink := 0; blink < iterations; blink++ {
		newCounts := make(map[int]int)

		for stone, count := range stoneCounts {
			if stone == 0 {
				newCounts[1] += count
				continue
			}

			digits := len(strconv.Itoa(stone))
			if digits & 1 == 1 {
				newCounts[stone * 2024] += count
				continue
			}

			half := digits/2
			div := int(math.Pow10(half))

			left := stone / div
			right := stone % div

			newCounts[left] += count
			newCounts[right] += count
		}

		stoneCounts = newCounts
	}

	result := 0
	for _, count := range stoneCounts {
		result += count
	}

	return result
}
