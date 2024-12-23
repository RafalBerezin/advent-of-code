package lib

var Dirs4 = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var Dirs4Diagonal = []Point{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
var Dirs8 = []Point{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

func ByteDir(character byte) Point {
	switch character {
	case '^':
		return Dirs4[0]
	case '>':
		return Dirs4[1]
	case 'v':
		return Dirs4[2]
	case '<':
		return Dirs4[3]
	default:
		return Point{0, 0}
	}
}

type Point struct {
	X, Y int
}

func (p *Point) Add(other *Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p *Point) Sub(other *Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}
