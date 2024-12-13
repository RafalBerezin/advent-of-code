package day13

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

var clawOffset = 10_000_000_000_000

func Part2(file *lib.InputFile) any {
	machines := *parseInput(file)
	for i := range machines {
		machines[i].prizeX += clawOffset
		machines[i].prizeY += clawOffset
	}

	result := 0
	for _, machine:= range machines {
		if cost := checkMachine2(&machine); cost != -1 {
			result += cost
		}
	}

	return result
}

func checkMachine2(pMachine *machine) int {
	machine := *pMachine

	c := machine.ax * machine.by - machine.ay * machine.bx
	if c == 0 {
		return -1
	}

	d := machine.prizeY * machine.ax - machine.prizeX * machine.ay
	if d % c != 0 {
		return -1
	}

	bPresses := d / c
	if (machine.prizeX - (bPresses * machine.bx)) % machine.ax != 0 {
		return -1
	}

	aPresses := (machine.prizeX - bPresses * machine.bx) / machine.ax
	if aPresses * machine.ax + bPresses * machine.bx != machine.prizeX ||
		aPresses * machine.ay + bPresses  * machine.by != machine.prizeY {
		return -1
	}

	return aPresses * ACost + bPresses * BCost
}
