package day_eight

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/louisandrew/advent-of-code/2024/utils"
)

type antennaMap struct {
	cr        utils.CoordinateRange
	antinodes []utils.Coordinate
}

func (am *antennaMap) addAntinode(c utils.Coordinate) {
	for _, a := range am.antinodes {
		if a[0] == c[0] && a[1] == c[1] {
			return
		}
	}

	am.antinodes = append(am.antinodes, c)
}

type freq []utils.Coordinate

func (f freq) getAntinodes(am *antennaMap) {
	if len(f) < 2 {
		return
	}

	for i := 0; i < len(f)-1; i++ {
		for j := i + 1; j < len(f); j++ {
			a := f[i]
			b := f[j]
			diff := b.Diff(a)

			am.addAntinode(a)
			am.addAntinode(b)

			nextPos1 := utils.Coordinate{a[0] - diff[0], a[1] - diff[1]}
			for am.cr.IsInBounds(nextPos1) {
				am.addAntinode(nextPos1)
				nextPos1 = utils.Coordinate{nextPos1[0] - diff[0], nextPos1[1] - diff[1]}
			}

			nextPos2 := utils.Coordinate{b[0] + diff[0], b[1] + diff[1]}
			for am.cr.IsInBounds(nextPos2) {
				am.addAntinode(nextPos2)
				nextPos2 = utils.Coordinate{nextPos2[0] + diff[0], nextPos2[1] + diff[1]}
			}
		}
	}

}

type freqMap map[string]freq
type antennas struct {
	freqMap freqMap
}

func parseMap(input string) (antennas, antennaMap) {
	lines := strings.Split(input, "\n")
	freqMap := make(map[string]freq)

	r, _ := regexp.Compile("[a-zA-Z0-9]")
	for y, line := range lines {
		matchIndices := r.FindAllStringIndex(line, -1)
		for _, matchIdx := range matchIndices {
			x := matchIdx[0]
			antenna := string(line[x])
			c := utils.Coordinate{x, y}

			if _, ok := freqMap[antenna]; !ok {
				freqMap[antenna] = []utils.Coordinate{c}
			} else {
				freqMap[antenna] = append(freqMap[antenna], c)
			}
		}
	}

	return antennas{freqMap}, antennaMap{
		cr: utils.CoordinateRange{Min: utils.Coordinate{0, 0}, Max: utils.Coordinate{len(lines[0]), len(lines)}},
	}
}

func Solution(input string) int {
	antennas, antennaMap := parseMap(input)

	for _, v := range antennas.freqMap {
		v.getAntinodes(&antennaMap)

	}
	fmt.Println(antennaMap.antinodes, antennaMap.cr)

	return len(antennaMap.antinodes)
}
