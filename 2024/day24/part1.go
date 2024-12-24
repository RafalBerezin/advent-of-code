package day24

import (
	"slices"
	"strconv"
	"strings"

	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

type gate struct {
	in1, in2, out string
	op byte
	processed bool
}

func Part1(file *lib.InputFile) any {
	input := file.Strings()
	splitIndex := slices.Index(input, "")

	valueDefinitions := input[:splitIndex]
	gateDefinitions := input[splitIndex+1:]

	values := make(map[string]bool)
	for _, value := range valueDefinitions {
		values[value[:3]] = value[5] == '1'
	}

	gates := make([]gate, 0, len(input)-splitIndex-1)
	for _, gateDef := range gateDefinitions {
		parts := strings.Split(gateDef, " ")
		gates = append(gates, gate{
			in1: parts[0], 
			in2: parts[2],
			out: parts[4],
			op: parts[1][0],
			processed: false,
		})
	}

	result := 0
	remainingGates := len(gates)

	for remainingGates > 0 {
		for i := range gates {
			gate := gates[i]
			if gate.processed {
				continue
			}

			v1, found := values[gate.in1]
			if !found {
				continue
			}

			v2, found := values[gate.in2]
			if !found {
				continue
			}

			var gateResult bool
			switch gate.op {
			case 'A':
				gateResult = v1 && v2
			case 'X':
				gateResult = v1 != v2
			case 'O':
				gateResult = v1 || v2
			}
			values[gate.out] = gateResult
			gates[i].processed = true
			remainingGates--

			if gateResult && gate.out[0] == 'z' {
				num, err := strconv.Atoi(gate.out[1:])
				lib.CheckError(err)
				result |= 1 << num
			}
		}
	}

	return result
}
