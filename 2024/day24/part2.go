package day24

import (
	"slices"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

/*
A simple Full Adder

X >─┬───┐   ┌───┐
    │   XOR─┤   XOR──> Z
    AND─┼───┼───┼───┐
Y >─┴───┘   AND─┼─┐ OR──> C
C >─────────┴───┘ └─┘

Z = X xor Y xor C
C = ((X xor Y) and C) or (X and Y)

Drawing this took longer than solving the problem
*/

type partialGate struct {
	in2, out string
}

func Part2(file *lib.InputFile) any {
	input := file.Strings()
	splitIndex := slices.Index(input, "")

	if splitIndex < 50 {
		return "Part 2 code not compatible with examples"
	}

	inputsCount := splitIndex / 2
	gateDefinitions := input[splitIndex+1:]

	inXors := make([]string, inputsCount)
	inAnds := make([]string, inputsCount)

	xors := make(map[string]partialGate)
	ands := make(map[string]partialGate)
	ors := make(map[string]partialGate)

	for _, gateDef := range gateDefinitions {
		parts := strings.Split(gateDef, " ")
		in1 := parts[0]
		op := parts[1][0]
		in2 := parts[2]
		out := parts[4]

		var opMap map[string]partialGate
		switch op {
		case 'A' :
			opMap = ands
		case 'X' :
			opMap = xors
		case 'O' :
			opMap = ors
		}

		if in1[0] == 'x' || in2[0] == 'x' {
			i := ((in1[1] - '0') * 10) + in1[2] - '0'
			if op == 'X' {
				inXors[i] = out
				} else {
				inAnds[i] = out
			}
			continue
		}

		opMap[in1] = partialGate{in2, out}
		opMap[in2] = partialGate{in1, out}
	}

	errors := make([]string, 0, 8)

	xorOut := inXors[0]
	andOut := inAnds[0]
	carry := andOut

	if xorOut != "z00" {
		errors = append(errors, xorOut)
		errors = append(errors, andOut)
		carry = xorOut
	}

	for i := 1; i < inputsCount; i++ {
		inXor := inXors[i]
		inAnd := inAnds[i]

		carryXor := xors[carry]
		if carryXor.in2 != inXor {
			errors = append(errors, inXor)
		}
		if carryXor.out[0] != 'z' {
			errors = append(errors, carryXor.out)
		}

		carryAnd := ands[carry].out

		or, found := ors[inAnd]
		if or.in2 != carryAnd {
			if found {
				errors = append(errors, carryAnd)
			} else {
				errors = append(errors, inAnd)
				or = ors[carryAnd]
			}
		}

		if len(errors) & 1 == 1 {
			carry = errors[len(errors)-1]
			errors = append(errors, or.out)
		} else {
			carry = or.out
		}
	}

	slices.Sort(errors)
	return strings.Join(errors, ",")
}
