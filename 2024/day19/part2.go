package day19

import (
	"strings"
	"sync"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	input := file.Strings()

	towels := strings.Split(input[0], ", ")
	designs := input[2:]

	wg := sync.WaitGroup{}
	mut := sync.Mutex{}

	wg.Add(len(designs))

	result := 0
	for _, design := range designs {
		go func() {
			count := countPossibleArrangements(design, &towels, make(map[string]int))

			if count > 0 {
				mut.Lock()
				result += count
				mut.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	return result
}

func countPossibleArrangements(design string, pTowels *[]string, cache map[string]int) int {
	dLen := len(design)
	if dLen == 0 {
		return 1
	}

	if cache[design] != 0 {
		return cache[design]
	}

	totalCount := 0

	towels := *pTowels
	for _, towel := range towels {
		tLen := len(towel)
		if tLen > dLen || design[:tLen] != towel {
			continue
		}

		if tLen == dLen {
			totalCount++
			continue
		}

		next := design[tLen:]

		if cache[next] == -1 {
			continue
		}

		count := countPossibleArrangements(design[tLen:], pTowels, cache)
		if count > 0 {
			totalCount += count
		}
	}

	if totalCount == 0 {
		totalCount = -1
	}

	cache[design] = totalCount
	return totalCount
}
