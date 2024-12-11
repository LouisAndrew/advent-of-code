package utils

type Direction [2]int

type Vector struct {
	D Direction
	C Coordinate
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
