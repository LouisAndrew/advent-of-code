package utils

type Range struct {
	Min int
	Max int
}

func (r Range) Length() int {
	return r.Max - r.Min
}
