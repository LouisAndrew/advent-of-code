package day_four

import (
	"regexp"
	"strings"

	"github.com/louisandrew/advent-of-code/2024/utils"
)

// Get all indices for X,M,A,S
// Iterate over all X indices, going through all directions possible

const (
	X = "X"
	M = "M"
	A = "A"
	S = "S"
)

var charMap = map[string][][]int{}

//	First part
//
// var charOrder = []string{X, M, A, S}
var charOrder = []string{M, A, S}

func getCharIndices(input string, char string) []int {
	r, _ := regexp.Compile(char)
	matches := r.FindAllStringIndex(input, -1)
	out := []int{}

	for _, match := range matches {
		out = append(out, match[0])
	}

	return out
}
func populateCharMap(input string, yCoordinate int) {
	for _, char := range charOrder {
		indices := getCharIndices(input, char)
		for _, index := range indices {
			charMap[char][yCoordinate] = append(charMap[char][yCoordinate], index)
		}
	}
}
func parseLines(input string) {
	lines := strings.Split(input, "\n")
	for _, char := range charOrder {
		charMap[char] = make([][]int, len(lines)+1)
	}

	for i, line := range lines {
		if line == "" {
			continue
		}

		populateCharMap(line, i)
	}
}

func charExistsAtCoordinate(x, y int, char string) bool {
	for _, indices := range charMap[char][y] {
		if x == indices {
			return true
		}
	}

	return false
}

// First part
// func findNextChar(x, y int, direction direction, charIndex int) bool {
// 	if charIndex == 3 {
// 		return true
// 	}
//
// 	char := charOrder[charIndex+1]
// 	if y+direction[1] > len(charMap[char]) || y+direction[1] < 0 {
// 		return false
// 	}
//
// 	if charExistsAtCoordinate(x+direction[0], y+direction[1], char) {
// 		return findNextChar(x+direction[0], y+direction[1], direction, charIndex+1)
// 	}
//
// 	return false
// }

func findCharInDirection(x, y int, direction utils.Direction, char string) bool {
	if y+direction[1] > len(charMap[char]) || y+direction[1] < 0 {
		return false
	}

	return charExistsAtCoordinate(x+direction[0], y+direction[1], char)
}

func checkForXShapedMAS(x, y int) bool {
	firstCondition := false
	secondCondition := false

	if findCharInDirection(x, y, utils.UpLeft, M) && findCharInDirection(x, y, utils.DownRight, S) {
		firstCondition = true
	} else if findCharInDirection(x, y, utils.UpLeft, S) && findCharInDirection(x, y, utils.DownRight, M) {
		firstCondition = true
	}

	if findCharInDirection(x, y, utils.UpRight, M) && findCharInDirection(x, y, utils.DownLeft, S) {
		secondCondition = true
	} else if findCharInDirection(x, y, utils.UpRight, S) && findCharInDirection(x, y, utils.DownLeft, M) {
		secondCondition = true
	}

	return firstCondition && secondCondition
}

func Solution(input string) int {
	parseLines(input)

	out := 0
	/// solution for the first part
	// for yCoordinate, xIndices := range charMap[X] {
	// 	for _, xIndex := range xIndices {
	// 		for _, direction := range [8]utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right, utils.UpLeft, utils.UpRight, utils.DownLeft, utils.DownRight} {
	// 			if findNextChar(xIndex, yCoordinate, direction, 0) {
	// 				out += 1
	// 			}
	// 		}
	// 	}
	// }
	/// Second part:
	for yCoordinate, xIndices := range charMap[A] {
		for _, xIndex := range xIndices {
			if checkForXShapedMAS(xIndex, yCoordinate) {
				out += 1
			}
		}
	}

	return out
}
