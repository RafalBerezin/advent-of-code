package day8

import (
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

func Part2(file *lib.InputFile) any {
	grid := file.ByteGrid()

	height := len(grid)
	width := len(grid[0])

	// can also use a map
	// antennas := make(map[byte][][]int)
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

				antinodes[first[0] * width + first[1]] = true
				antinodes[second[0] * width + second[1]] = true

				beforeRow := first[0]
				beforeCol := first[1]
				for {
					beforeRow -= dRow
					beforeCol -= dCol

					if beforeRow < 0 || beforeRow >= height || beforeCol < 0 || beforeCol >= width {
						break
					}

					antinodes[beforeRow * width + beforeCol] = true
					grid[beforeRow][beforeCol] = '#'
				}


				afterRow := second[0]
				afterCol := second[1]
				for {
					afterRow += dRow
					afterCol += dCol

					if afterRow < 0 || afterRow >= height || afterCol < 0 || afterCol >= width {
						break
					}

					antinodes[afterRow * width + afterCol] = true
					grid[afterRow][afterCol] = '#'
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
