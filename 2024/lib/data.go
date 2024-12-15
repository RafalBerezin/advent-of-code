package lib

import "slices"

var Dirs4 = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var Dirs4Diagonal = [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
var Dirs8 = [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

var byteDirs = []byte{'^', '>', 'v', '<'}

func ByteDir(character byte) []int {
	i := slices.Index(byteDirs, character)
	if i == -1 {
		return []int{}
	}
	return Dirs4[i]
}
