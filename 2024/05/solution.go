package day_five

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type rule string
type ruleset string

func (rs *ruleset) addRule(r rule) {
	if *rs == "" {
		*rs = ruleset(r)
		return
	}

	*rs = ruleset(*rs + "," + ruleset(r))
}

func (rs *ruleset) containsOrderingRule(x, y int) bool {
	s := fmt.Sprintf(`%d\|%d`, x, y)
	r, _ := regexp.Compile(s)
	return r.MatchString(string(*rs))
}

type sequence []int

func (s *sequence) sort(rs ruleset) sequence {
	copySeq := make(sequence, len(*s))
	copy(copySeq, *s)
	n := len(copySeq)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if rs.containsOrderingRule((copySeq)[j+1], (copySeq)[j]) {
				// Swap elements
				(copySeq)[j], (copySeq)[j+1] = (copySeq)[j+1], (copySeq)[j]
			}
		}
	}

	return copySeq
}

const RULE_SEPARATOR = "|"

func parseInput(input string) (ruleset, []sequence) {
	sections := strings.Split(input, "\n\n")
	rs := ruleset("")
	sequences := []sequence{}

	for _, ruleStr := range strings.Split(sections[0], "\n") {
		if strings.Contains(ruleStr, RULE_SEPARATOR) {
			rs.addRule(rule(ruleStr))
		}
	}

	for _, updateStr := range strings.Split(sections[1], "\n") {
		seq := sequence{}
		for _, pageStr := range strings.Split(updateStr, ",") {
			pageNum, _ := strconv.Atoi(pageStr)
			seq = append(seq, pageNum)
		}

		sequences = append(sequences, seq)
	}

	return rs, sequences
}

func Solution(input string) int {
	rs, sequences := parseInput(input)

	out := 0
	for _, seq := range sequences {
		sortedSeq := seq.sort(rs)
		// First part
		// if reflect.DeepEqual(seq, sortedSeq) {
		// 	middleIndex := len(seq) / 2
		// 	out += seq[middleIndex]
		// }

		// Second part
		if !reflect.DeepEqual(seq, sortedSeq) {
			middleIndex := len(seq) / 2
			out += sortedSeq[middleIndex]
		}
	}

	return out
}
