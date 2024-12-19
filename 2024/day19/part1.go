package day19

import (
	"strings"
	"sync"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Strings()

	towels := strings.Split(input[0], ", ")
	designs := input[2:]

	wg := sync.WaitGroup{}
	mut := sync.Mutex{}

	wg.Add(len(designs))

	result := 0
	for _, design := range designs {
		go func() {
			found := checkDesign(design, &towels, make(map[string]int))

			if found {
				mut.Lock()
				result++
				mut.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	return result
}

func checkDesign(design string, pTowels *[]string, cache map[string]int) bool {
	dLen := len(design)
	if dLen == 0 {
		cache[design] = 1
		return true
	}

	if cache[design] != 0 {
		return cache[design] == 1
	}

	towels := *pTowels
	for _, towel := range towels {
		tLen := len(towel)
		if tLen > dLen || design[:tLen] != towel {
			continue
		}

		found := checkDesign(design[tLen:], pTowels, cache)
		if found {
			cache[design] = 1
			return true
		}
	}

	cache[design] = -1
	return false
}
