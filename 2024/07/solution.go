package day_seven

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	ADD    = "+"
	MUL    = "*"
	CONCAT = "||"
)

var operators = []string{ADD, MUL, CONCAT}
var operatorMap = map[string]func(int, int) int{
	ADD: func(a, b int) int { return a + b },
	MUL: func(a, b int) int { return a * b },
	CONCAT: func(a, b int) int {
		aStr := strconv.Itoa(a)
		bStr := strconv.Itoa(b)
		abStr := aStr + bStr
		ab, _ := strconv.Atoi(abStr)
		return ab
	},
}

type equation struct {
	result    int
	nums      []int
	operators []string
}

func parseLines(input string) []equation {
	equations := []equation{}
	for _, line := range strings.Split(input, "\n") {
		e := equation{}

		r, _ := regexp.Compile(`\d+`)
		matches := r.FindAllString(line, -1)
		result, _ := strconv.Atoi(matches[0])

		e.result = result

		for _, match := range matches[1:] {
			num, _ := strconv.Atoi(match)
			e.nums = append(e.nums, int(num))
		}

		equations = append(equations, e)
	}

	return equations
}

func evaluate(stack []int) []int {
	if len(stack) == 1 {
		return stack
	}

	a, b, rest := stack[0], stack[1], stack[2:]
	results := []int{}
	for _, op := range operators {
		result := operatorMap[op](a, b)
		results = append(results, evaluate(append([]int{result}, rest...))...)
	}

	return results
}

func Solution(input string) int {
	e := parseLines(input)
	out := 0
	for _, eq := range e {
		results := evaluate(eq.nums)
		for _, r := range results {
			if eq.result == r {
				out += eq.result
				break
			}
		}
	}

	return out
}
