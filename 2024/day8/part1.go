package day8

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part1(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	antennas := make([][][]int, 255)

	for i, row := range grid {
		for j, col := range row {
			if col == '.' {
				continue
			}

			antennas[col] = append(antennas[col], []int{i, j})
		}
	}

	antinodes := make([]bool, height * width)

	for _, frequency := range antennas {
		for i := 0 ; i < len(frequency)-1; i++ {
			first := frequency[i]
			for j := i + 1; j < len(frequency); j++ {
				second := frequency[j]
				dRow := second[0] - first[0]
				dCol := second[1] - first[1]

				beforeRow := first[0] - dRow
				beforeCol := first[1] - dCol
				if beforeRow >= 0 && beforeRow < height && beforeCol >= 0 && beforeCol < width {
					antinodes[beforeRow * width + beforeCol] = true
				}

				afterRow := second[0] + dRow
				afterCol := second[1] + dCol
				if afterRow >= 0 && afterRow < height && afterCol >= 0 && afterCol < width {
					antinodes[afterRow * width + afterCol] = true
				}
			}
		}
	}

	result := 0
	for _, antinode := range antinodes {
		if antinode {
			result ++
		}
	}

	return result
}
