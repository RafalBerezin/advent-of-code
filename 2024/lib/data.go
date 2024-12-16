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

type Point struct {
	X, Y int
}

func (p *Point) Add(other *Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}
