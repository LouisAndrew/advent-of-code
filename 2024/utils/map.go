package utils

type Map struct {
	cr CoordinateRange
}

func (m *Map) SetMin(c Coordinate) {
	(*m).cr.Min = c
}

func (m *Map) SetMax(c Coordinate) {
	(*m).cr.Max = c
}

func (m Map) IsInBounds(c Coordinate) bool {
	x, y := c[0], c[1]
	return x >= m.cr.Min[0] && y >= m.cr.Min[1] && x < m.cr.Max[0] && y < m.cr.Max[1]
}
