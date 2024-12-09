package day_six

import (
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/louisandrew/advent-of-code/2024/utils"
)

type patrolMap struct {
	cr        utils.CoordinateRange
	obstacles []utils.Coordinate
	marked    []utils.Vector
}

func (pm *patrolMap) isCoordinateMarked(v utils.Vector) (bool, int) {
	for i, mark := range pm.marked {
		if mark.C[0] == v.C[0] && mark.C[1] == v.C[1] {
			return true, i
		}
	}

	return false, -1
}
func (pm *patrolMap) hasBeenHere(v utils.Vector) bool {
	isMarked, idx := pm.isCoordinateMarked(v)
	if !isMarked {
		return false
	}

	return reflect.DeepEqual(pm.marked[idx].D, v.D)
}
func (pm *patrolMap) mark(v utils.Vector) {
	isMarked, idx := pm.isCoordinateMarked(v)

	if isMarked {
		pm.marked[idx].D = v.D
		return
	}

	pm.marked = append(pm.marked, v)
}
func (pm patrolMap) isObstacle(c utils.Coordinate) bool {
	for _, obs := range pm.obstacles {
		obx, oby := obs[0], obs[1]
		if c[0] == obx && c[1] == oby {
			return true
		}
	}

	return false
}

type guard struct {
	d        utils.Direction
	pos      utils.Coordinate
	isInLoop bool
}

func (g *guard) turn() {
	switch g.d {
	case utils.Up:
		g.d = utils.Right
	case utils.Right:
		g.d = utils.Down
	case utils.Down:
		g.d = utils.Left
	case utils.Left:
		g.d = utils.Up
	}
}
func (g *guard) walk(patrolMap *patrolMap) bool {
	newPos := g.pos.Add(g.d)

	if patrolMap.hasBeenHere(utils.Vector{D: g.d, C: g.pos}) {
		g.isInLoop = true
		return false
	}

	if !patrolMap.cr.IsInBounds(newPos) {
		// Out of bound. what to do
		patrolMap.mark(utils.Vector{D: g.d, C: g.pos})
		return false
	}
	if patrolMap.isObstacle(newPos) {
		g.turn()
		return true
	}

	patrolMap.mark(utils.Vector{D: g.d, C: g.pos})
	g.pos = newPos
	return true
}

func parsePatrolMap(input string) (patrolMap, guard) {
	guard := guard{}
	obstacles := []utils.Coordinate{}

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}

		guardRegex, _ := regexp.Compile(`\^|<|>|v`)
		guardMatchIndex := guardRegex.FindStringIndex(line)
		if guardMatchIndex != nil {
			guardMatch := line[guardMatchIndex[0]]
			guard.pos = utils.Coordinate{guardMatchIndex[0], i}
			switch string(guardMatch) {
			case "^":
				guard.d = utils.Up
			case ">":
				guard.d = utils.Right
			case "v":
				guard.d = utils.Down
			case "<":
				guard.d = utils.Left
			}
		}

		obstacleRegex, _ := regexp.Compile("#")
		obstacleMatches := obstacleRegex.FindAllStringIndex(line, -1)
		for _, match := range obstacleMatches {
			obstacles = append(obstacles, utils.Coordinate{match[0], i})
		}
	}

	patrolMap := patrolMap{utils.CoordinateRange{Min: utils.Coordinate{0, 0}, Max: utils.Coordinate{len(lines[0]), len(lines)}}, obstacles, []utils.Vector{}}

	return patrolMap, guard
}

const DEBUG = true

// Part one
// func Solution(input string) int {
// 	patrolMap, guard := parsePatrolMap(input)
// 	patrolMap.obstacles = append(patrolMap.obstacles, utils.Coordinate{1, 8})
//
// 	for guard.walk(&patrolMap) {
// 	}
//
// 	return len(patrolMap.marked)
// }

func Solution(input string) int {
	patrolMap, g := parsePatrolMap(input)

	originalPatrolMap := patrolMap
	originalGuard := g

	for g.walk(&patrolMap) {
	}

	out := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, m := range patrolMap.marked {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newGuard := originalGuard
			newMap := originalPatrolMap

			// Concurrency := make sure you are not modifying the same data...
			newMap.obstacles = append([]utils.Coordinate{}, newMap.obstacles...)
			newMap.marked = append([]utils.Vector{}, newMap.marked...)
			newMap.obstacles = append(newMap.obstacles, m.C)

			for newGuard.walk(&newMap) {
			}

			if newGuard.isInLoop {
				mu.Lock()
				out++
				defer mu.Unlock()
			}
		}()
	}

	wg.Wait()

	return out
}
