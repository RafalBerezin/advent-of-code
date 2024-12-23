package day22

const PRUNE = 16777216 -1 // -1 since i use & instead of %
const NUMBERS_PER_DAY = 2000

func nextSecret(secret int) int {
	secret = (secret ^ (secret << 6)) & PRUNE
	secret = (secret ^ (secret >> 5)) & PRUNE
	return (secret ^ (secret << 11)) & PRUNE
}
