package day21

import (
	"math"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

const (
	ROBOTS_1 = 2
	ROBOTS_2 = 25
)

type keypad map[byte]lib.Point

type cacheKey struct {
	sequence string
	depth int
}

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+

var doorKeypad = keypad{
	'7': {Y: 0, X: 0},
	'8': {Y: 0, X: 1},
	'9': {Y: 0, X: 2},
	'4': {Y: 1, X: 0},
	'5': {Y: 1, X: 1},
	'6': {Y: 1, X: 2},
	'1': {Y: 2, X: 0},
	'2': {Y: 2, X: 1},
	'3': {Y: 2, X: 2},
	' ': {Y: 3, X: 0},
	'0': {Y: 3, X: 1},
	'A': {Y: 3, X: 2},
}
var doorStart = doorKeypad['A']
var doorAvoid = doorKeypad[' ']

// +---+---+
// | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

var robotKeypad = keypad{
	' ': {Y: 0, X: 0},
	'^': {Y: 0, X: 1},
	'A': {Y: 0, X: 2},
	'<': {Y: 1, X: 0},
	'v': {Y: 1, X: 1},
	'>': {Y: 1, X: 2},
}
var robotStart = robotKeypad['A']
var robotAvoid = robotKeypad[' ']

func Part1(file *lib.InputFile) any {
	return getComplexities(file, ROBOTS_1, map[cacheKey]int{})
}

func Part2(file *lib.InputFile) any {
	return getComplexities(file, ROBOTS_2, map[cacheKey]int{})
}

func getComplexities(file *lib.InputFile, robots int, cache map[cacheKey]int) int {
	input := file.ByteGrid()
	result := 0

	for _, keys := range input {
		num, err := strconv.Atoi(string(keys[:len(keys)-1]))
		lib.CheckError(err)

		length := getKeypadSequenceLength(keys, doorStart, doorAvoid, doorKeypad, robots, cache)
		result += length * num
	}

	return result
}

func getKeypadSequenceLength(keys []byte, from, avoid lib.Point, keypad keypad, depth int, cache map[cacheKey]int) int {
	cacheKey := cacheKey{string(keys), depth}
	if cached, found := cache[cacheKey]; found {
		return cached
	}
	length := 0

	for _, key := range keys {
		to := keypad[key]
		if from == to {
			length++
			continue
		}

		sequences := getMoveSequences(from, to, avoid)

		if depth == 0 {
			length += len(sequences[0])
			from = to
			continue
		}

		minLength := math.MaxInt
		for _, sequence := range sequences {
			sequenceLength := getKeypadSequenceLength(sequence, robotStart, robotAvoid, robotKeypad, depth-1, cache)
			if sequenceLength < minLength {
				minLength = sequenceLength
			}
		}

		length += minLength
		from = to
	}

	if length != 0 {
		cache[cacheKey] = length
	}
	return length
}

func getMoveSequences(from, to, avoid lib.Point) [][]byte {
	delta := to.Sub(&from)

	colMovement := ">"
	if delta.X < 0 {
		colMovement = "<"
		delta.X = -delta.X
	}
	colMovement = strings.Repeat(colMovement, delta.X)

	rowMovement := "v"
	if delta.Y < 0 {
		rowMovement = "^"
		delta.Y = -delta.Y
	}
	rowMovement = strings.Repeat(rowMovement, delta.Y)

	colFirst := []byte(colMovement + rowMovement + "A")
	rowFirst := []byte(rowMovement + colMovement + "A")

	if from.X == avoid.X && to.Y == avoid.Y {
		return [][]byte{colFirst}
	}

	if from.Y == avoid.Y && to.X == avoid.X {
		return [][]byte{rowFirst}
	}

	return [][]byte{rowFirst, colFirst}
}
