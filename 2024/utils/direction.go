package utils

type Direction [2]int

func (d Direction) Add(d2 Direction) Direction {
	return Direction{d[0] + d2[0], d[1] + d2[1]}
}

func (d Direction) String() string {
	if d == Up {
		return "Up"
	}

	if d == Down {
		return "Down"
	}

	if d == Left {
		return "Left"
	}

	if d == Right {
		return "Right"
	}

	return ""
}

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
