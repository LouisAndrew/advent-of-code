package day_twelve

import (
	"fmt"
	"strings"

	"github.com/louisandrew/advent-of-code/2024/utils"
)

type node struct {
	pos utils.Coordinate
}

type region struct {
	f   *utils.CoordinateMap
	per int
	id  int
}

func (r region) String() string {
	out := ""
	for k, x := range *r.f {
		out += fmt.Sprintf("%d:%v,", k, x)
	}

	out += fmt.Sprintf("Perimeter: %d", r.per)

	return out
}

type garden struct {
	plants map[string]*utils.CoordinateMap
}

func parse(input string) garden {
	lines := strings.Split(input, "\n")
	g := garden{
		plants: make(map[string]*utils.CoordinateMap),
	}

	for y, line := range lines {
		for x, char := range line {
			plant := string(char)

			if _, ok := g.plants[plant]; !ok {
				g.plants[plant] = &utils.CoordinateMap{}
			}

			g.plants[plant].Add(utils.Coordinate{x, y})
		}
	}

	return g
}

var directions = []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

const MAX_PERIMETER = 4

func b(m *utils.CoordinateMap) [](*region) {
	regions := [](*region){}

	for x, v := range *m {
		for _, y := range v {
			c := utils.Coordinate{x, y}
			plantPer := 0
			var reg *region

			mergeRegions := [](*region){}
			dMap := map[utils.Direction]bool{
				utils.Up:    false,
				utils.Down:  false,
				utils.Left:  false,
				utils.Right: false,
			}

			for _, d := range directions {
				adj := c.Add(d)

				if !m.Has(adj) {
					continue
				}

				dMap[d] = true
				for _, r := range regions {
					if r.f.Has(adj) {
						if reg != nil {
							mergeRegions = append(mergeRegions, r)
							continue
						}

						reg = r
						break
					}
				}
			}

			// nums of sides == nums of corners (incl. inverted corner)
			for _, d := range []utils.Direction{utils.Up, utils.Down} {
				for _, d2 := range []utils.Direction{utils.Left, utils.Right} {
					if !dMap[d] && !dMap[d2] {
						plantPer++
						continue
					}

					// inverted corner
					if dMap[d] && dMap[d2] && !m.Has(c.Add(d).Add(d2)) {
						plantPer++
					}
				}
			}

			// plant has no region
			if reg == nil {
				cm := utils.CoordinateMap{}
				newRegion := &region{
					f: &cm,
				}

				reg = newRegion
				regions = append(regions, newRegion)
			}

			b := [](*region){}
			if len(mergeRegions) > 0 {
				for _, r := range regions {
					if r == reg {
						continue
					}

					merge := false
					for _, mr := range mergeRegions {
						if r == mr {
							merge = true
							break
						}
					}

					if !merge {
						b = append(b, r)
						continue
					}

					reg.f.Merge(r.f)
					reg.per += r.per
				}
			}

			reg.f.Add(c)
			reg.per += plantPer
		}
	}

	return regions
}

func Solution(input string) int {
	// parse all plots, store positions in a map
	// for each plant type, build a region
	g := parse(input)
	out := 0
	for _, m := range g.plants {
		regions := b(m)

		for _, r := range regions {
			area := r.f.Length()
			per := r.per

			out += area * per
		}
	}

	return out
}
