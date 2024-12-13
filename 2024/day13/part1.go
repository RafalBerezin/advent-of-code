package day13

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	machines := *parseInput(file)

	result := 0
	for _, machine:= range machines {
		cost := checkMachine(&machine)
		if cost != -1 {
			result += cost
		}
	}

	return result
}

func checkMachine(pMachine *machine) int {
	machine := *pMachine

	lowestCost := -1

	for aPresses := 0; aPresses <= 100; aPresses++ {
		for bPresses := 0; bPresses <= 100; bPresses++ {
			posX := aPresses * machine.ax + bPresses * machine.bx
			posY := aPresses * machine.ay + bPresses * machine.by

			if posX == machine.prizeX && posY == machine.prizeY {
				cost := aPresses * ACost + bPresses * BCost
				if lowestCost == -1 || cost < lowestCost {
					lowestCost = cost
				}
				break
			}
		}
	}

	return lowestCost
}
