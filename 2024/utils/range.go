package utils

type R[V int | Coordinate] struct {
	Min V
	Max V
}

type Range R[int]

func (r Range) Length() int {
	return r.Max - r.Min
}

type CoordinateRange R[Coordinate]

func (cr CoordinateRange) IsInBounds(c Coordinate) bool {
	return c[0] >= cr.Min[0] && c[1] >= cr.Min[1] && c[0] < cr.Max[0] && c[1] < cr.Max[1]
}
