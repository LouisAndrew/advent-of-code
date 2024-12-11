package utils

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

// X, Y Map
type CoordinateMap map[int][]int

func (m CoordinateMap) Has(c Coordinate) bool {
	x, y := c[0], c[1]
	if _, ok := m[x]; !ok {
		return false
	}

	for _, i := range m[x] {
		if i == y {
			return true
		}
	}

	return false
}

func (m *CoordinateMap) Add(c Coordinate) {
	if _, ok := (*m)[c[0]]; !ok {
		(*m)[c[0]] = []int{}
	}

	(*m)[c[0]] = append((*m)[c[0]], c[1])
}

func BuildCoordinateMap(coordinates []Coordinate) CoordinateMap {
	m := CoordinateMap{}
	for _, c := range coordinates {
		m.Add(c)
	}

	return m
}
