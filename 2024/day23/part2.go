package day23

import (
	"slices"
	"strings"
	"sync"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
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

	biggestNetwork := []string{"No network found"}
	biggestSize := 0

	wg := sync.WaitGroup{}
	mut := sync.Mutex{}

	for triplet := range triplets {
		wg.Add(1)

		go func() {
			network := strings.Split(triplet, ",")

			for _, candidate := range connections[network[0]] {
				found := true
				for _, member := range network[1:] {
					if !slices.Contains(connections[member], candidate) {
						found = false
						break
					}
				}

				if found {
					network = append(network, candidate)
				}
			}
			networkSize := len(network)

			if networkSize >= biggestSize {
				mut.Lock()
				biggestSize = networkSize
				biggestNetwork = network
				mut.Unlock()
			}

			wg.Done()
		}()
	}

	slices.Sort(biggestNetwork)
	return strings.Join(biggestNetwork, ",")
}
