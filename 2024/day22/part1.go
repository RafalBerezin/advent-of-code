package day22

import (
	"sync"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Ints()

	wg := sync.WaitGroup{}
	wg.Add(len(input))

	mut := sync.Mutex{}
	result := 0

	for _, number := range input {
		go func() {
			secret := number
			for i := 0; i < NUMBERS_PER_DAY; i++ {
				secret = nextSecret(secret)
			}

			mut.Lock()
			result += secret
			mut.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	return result
}
