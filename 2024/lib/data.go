package lib

import (
	"slices"
)

var Dirs4 = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var Dirs4Diagonal = []Point{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
var Dirs8 = []Point{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

var byteDirs = []byte{'^', '>', 'v', '<'}

func ByteDir(character byte) Point {
	i := slices.Index(byteDirs, character)
	if i == -1 {
		return Point{0, 0}
	}
	return Dirs4[i]
}

type Point struct {
	X, Y int
}

func (p *Point) Add(other *Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}
