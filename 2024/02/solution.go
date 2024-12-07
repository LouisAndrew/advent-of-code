package day_two

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type level int
type report []level

func parseReports(input string) []report {
	reports := []report{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		r, _ := regexp.Compile(`\d+`)
		matches := r.FindAllString(line, -1)

		levels := []level{}
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}

			levels = append(levels, level(num))
		}

		reports = append(reports, levels)
	}

	return reports
}

const (
	NONE     = "NONE"
	ASC      = "ASC"
	DESC     = "DESC"
	MIN_DIFF = 1
	MAX_DIFF = 3
)

func (r report) isSave() bool {
	status := NONE
	for i := 1; i < len(r); i++ {
		curr := NONE

		a := r[i]
		b := r[i-1]
		if a == b {
			return false
		}

		if a > b {
			curr = ASC
		} else {
			curr = DESC
		}

		if status != NONE && status != curr {
			return false
		}

		diff := math.Abs(float64(a - b))
		if diff < MIN_DIFF || diff > MAX_DIFF {
			return false
		}

		status = curr
	}

	return true
}

func (r report) isSaveWithDampener() bool {
	isSave := r.isSave()
	if isSave {
		return true
	}

	for i := range len(r) {
		nr := report{}
		nr = append(nr, r[:i]...)
		nr = append(nr, r[i+1:]...)

		isSave = nr.isSave()
		if isSave {
			return true
		}
	}

	return false
}

func Solution(input string) int {
	reports := parseReports(input)
	valid := 0
	for _, r := range reports {
		if r.isSaveWithDampener() {
			valid++
		}
	}

	return valid
}
