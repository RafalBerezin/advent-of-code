package day23

import (
	"slices"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	input := file.Strings()

	connections := make(map[string][]string)
	for _, connection := range input {
		first := connection[:2]
		second := connection[3:]

		connections[first] = append(connections[first], second)
		connections[second] = append(connections[second], first)
	}

	triplets := make(map[string]bool)

	for from1, to1 := range connections {
		if from1[0] != 't' {
			continue
		}

		for _, from2 := range to1 {
			to2 := connections[from2]

			for _, from3 := range to2 {
				to3 := connections[from3]

				if slices.Contains(to3, from1) {
					triplet := []string{from1, from2, from3}
					slices.Sort(triplet)
					triplets[strings.Join(triplet, ",")] = true
				}
			}
		}
	}

	return len(triplets)
}
