package day_ten

import (
	"strconv"
	"strings"

	"github.com/louisandrew/advent-of-code/2024/utils"
)

const (
	START_POINT = 0
	MAX_HEIGHT  = 9
)

func parse(input string) topographicMap {
	lines := strings.Split(input, "\n")
	tm := topographicMap{
		m: utils.Map{},
	}

	tm.m.SetMin(utils.Coordinate{0, 0})
	tm.m.SetMax(utils.Coordinate{len(lines[0]), len(lines)})

	trailMap := make(map[int][]utils.Coordinate)
	for y, line := range lines {
		for x, ch := range line {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				num = -1
			}

			if _, ok := trailMap[num]; !ok {
				trailMap[num] = []utils.Coordinate{}
			}

			trailMap[num] = append(trailMap[num], utils.Coordinate{x, y})
		}
	}

	tm.t = make(map[int]utils.CoordinateMap)
	for k, v := range trailMap {
		tm.t[k] = utils.BuildCoordinateMap(v)
	}

	return tm
}

type trail struct {
	height int
	c      utils.Coordinate
}

type topographicMap struct {
	m       utils.Map
	t       map[int]utils.CoordinateMap
	visited utils.CoordinateMap
}

func (tm topographicMap) hasTrail(t trail) bool {
	return tm.t[t.height].Has(t.c)
}
func (tm topographicMap) visit(t trail) bool {
	if tm.visited.Has(t.c) {
		return false
	}

	tm.visited.Add(t.c)
	return true
}

var directions = []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

func move(t trail, tm *topographicMap) int {
	if t.height == MAX_HEIGHT {
		return 1
	}

	next := t.height + 1
	out := 0
	for _, d := range directions {
		n := t.c.Add(d)
		t := trail{height: next, c: n}
		// solution 1 : add `&& tm.visit(t)`
		if tm.hasTrail(t) {
			x := move(t, tm)
			out += x
		}
	}

	return out
}

func Solution(input string) int {
	tm := parse(input)

	o := 0
	for x, yArr := range tm.t[0] {
		for _, y := range yArr {
			// get a copy
			t := tm
			t.visited = make(utils.CoordinateMap)
			trail := trail{height: 0, c: utils.Coordinate{x, y}}
			x := move(trail, &t)
			o += x
		}
	}

	return o
}
