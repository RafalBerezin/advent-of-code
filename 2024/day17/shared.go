package day17

import (
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

const (
	ADV = 0
	BXL = 1
	BST = 2
	JNZ = 3
	BXC = 4
	OUT = 5
	BDV = 6
	CDV = 7
)

func parseInput(file *lib.InputFile) ([]int, []byte) {
	lines := file.Strings()

	registers := make([]int, 3)
	for reg, line := range lines[:3] {
		value, err := strconv.Atoi(line[12:])
		lib.CheckError(err)
		registers[reg] = value
	}

	iStrs := strings.Split(lines[4:5][0][9:], ",")
	instructions := make([]byte, len(iStrs))
	for i, instruction := range iStrs {
		instructions[i] = byte(instruction[0]) - '0'
	}

	return registers, instructions
}

func runProgram(pRegisters *[]int, pInstructions *[]byte) []byte {
	registers := *pRegisters
	instructions := *pInstructions

	out := make([]byte, 0)

	for i := 0; i < len(instructions); i += 2 {
		opcode := instructions[i]
		literal := instructions[i+1]

		switch opcode {
		case ADV, BDV, CDV:
			res := registers[0] >> getComboOperand(literal, &registers)

			switch opcode {
			case ADV:
				registers[0] = res
			case BDV:
				registers[1] = res
			case CDV:
				registers[2] = res
			}

		case BXL:
			registers[1] ^= int(literal)

		case BST:
			registers[1] = getComboOperand(literal, &registers) & 7

		case JNZ:
			if registers[0] != 0 {
				i = int(literal) - 2
			}

		case BXC:
			registers[1] ^= registers[2]

		case OUT:
			out = append(out, byte(getComboOperand(literal, &registers) & 7))
		}
	}

	return out
}

func getComboOperand(operand byte, registers *[]int) int {
	if operand <= 3 {
		return int(operand)
	}

	if operand == 7 {
		panic("combo operand 7: invalid program")
	}

	return (*registers)[operand - 4]
}
