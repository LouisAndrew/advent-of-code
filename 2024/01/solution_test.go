package day_one

import (
	"reflect"
	"testing"
)

const input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestParseIDList(t *testing.T) {
	got := parseIdList(input)
	want := [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
