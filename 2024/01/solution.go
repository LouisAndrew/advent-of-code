package day_one

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseIdList(input string) [2][]int {
	output := [2][]int{{}, {}}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		r, _ := regexp.Compile(`\d+`)
		matches := r.FindAllString(line, -1)
		if len(matches) != 2 {
			fmt.Println("Error lengths: ", matches, len(matches))
			continue
		}

		for i, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				fmt.Println("Error conversion: ", err)
				continue
			}

			output[i] = append(output[i], num)
		}
	}

	return output
}

func Solution(input string) string {
	output := 0

	similarityScoreMap := map[int]int{}
	idList := parseIdList(input)

	for _, id := range idList[1] {
		_, ok := similarityScoreMap[id]
		if !ok {
			similarityScoreMap[id] = 1
			continue
		}

		similarityScoreMap[id] += 1
	}

	for _, id := range idList[0] {
		storedSimilarityScore, ok := similarityScoreMap[id]
		if !ok {
			continue
		}

		output += storedSimilarityScore * id
	}

	return strconv.Itoa(output)
}
