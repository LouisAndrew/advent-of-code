package utils

type Direction [2]int
type Coordinate [2]int

func (c Coordinate) Add(d Direction) Coordinate {
	return Coordinate{c[0] + d[0], c[1] + d[1]}
}

// Difference of two coordinates, in absolute
func (c Coordinate) Diff(c2 Coordinate) Coordinate {
	return Coordinate{
		c[0] - c2[0],
		c[1] - c2[1],
	}
}

type CoordinateRange struct {
	Min Coordinate
	Max Coordinate
}

func (cr CoordinateRange) IsInBounds(c Coordinate) bool {
	return c[0] >= cr.Min[0] && c[1] >= cr.Min[1] && c[0] < cr.Max[0] && c[1] < cr.Max[1]
}

type Vector struct {
	D Direction
	C Coordinate
}

type Map [][]int

func (m Map) IsInBounds(c Coordinate) bool {
	x, y := c[0], c[1]
	return x >= 0 && y >= 0 && x < len(m[0]) && y < len(m)
}

var (
	Up        = Direction{0, -1}
	Down      = Direction{0, 1}
	Left      = Direction{-1, 0}
	Right     = Direction{1, 0}
	UpLeft    = Direction{-1, -1}
	UpRight   = Direction{1, -1}
	DownLeft  = Direction{-1, 1}
	DownRight = Direction{1, 1}
)
