package day22

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

type sequence struct {
	a, b, c, d byte
}

func Part2(file *lib.InputFile) any {
	input := file.Ints()

	bestResult := 0
	currSeq := sequence{}
	results := make(map[sequence]int)

	for _, number := range input {
		seenSequence := make(map[sequence]bool)
		secret := number
		previousPrice := 0
		for i := 0; i < NUMBERS_PER_DAY; i++ {
			secret = nextSecret(secret)

			price := secret % 10
			currSeq = sequence{byte(price - previousPrice), currSeq.a, currSeq.b, currSeq.c}
			previousPrice = price

			if i < 4 || seenSequence[currSeq] {
				continue
			}
			seenSequence[currSeq] = true

			results[currSeq] += price
			if results[currSeq] > bestResult {
				bestResult = results[currSeq]
			}
		}
	}

	return bestResult
}
