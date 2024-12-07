package day_three

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	DO      = "do()"
	DONT    = "don't()"
	ENABLE  = 1
	DISABLE = 0
	DEBUG   = false
)

func parseMuls(line string, flag int) ([]string, int) {
	out := []string{}
	r, _ := regexp.Compile(`do\(\)|don\'t\(\)|mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(line, -1)

	if DEBUG {
		fmt.Println(matches)
	}

	for _, match := range matches {
		if DEBUG {
			fmt.Println(match, flag)
		}

		if match == DO {
			flag = ENABLE
			continue
		}
		if match == DONT {
			flag = DISABLE
			continue
		}

		if flag == DISABLE {
			continue
		}

		out = append(out, match)
	}

	return out, flag
}

func parseLines(input string) []string {
	m := []string{}
	flag := ENABLE
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		muls, newFlag := parseMuls(line, flag)
		flag = newFlag
		m = append(m, muls...)
	}

	return m
}

func multiplyMulArgs(mul string) int {
	r, _ := regexp.Compile(`\d{1,3}`)
	matches := r.FindAllString(mul, -1)

	if len(matches) != 2 {
		fmt.Println("Error lengths: ", matches, len(matches))
		return 0
	}

	num1, _ := strconv.Atoi(matches[0])
	num2, _ := strconv.Atoi(matches[1])

	return num1 * num2
}

func Solution(input string) int {
	out := 0
	m := parseLines(input)
	for _, mul := range m {
		out += multiplyMulArgs(mul)
	}

	return out
}
