package day_two

import (
	"reflect"
	"testing"
)

func TestParseReports(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9`
	got := parseReports(input)
	want := []report{[]level{7, 6, 4, 2, 1}, []level{1, 2, 7, 8, 9}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIsValid(t *testing.T) {
	testCases := []struct {
		r    report
		safe bool
	}{
		{
			report{7, 6, 4, 2, 1},
			true,
		},
		{
			report{1, 2, 7, 8, 9},
			false,
		},
		{
			report{8, 6, 4, 4, 1},
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run("Is valid", func(t *testing.T) {
			got := testCase.r.isSave()
			want := testCase.safe
			if got != want {
				t.Errorf("got %v want %v; data %v", got, want, testCase.r)
			}
		})
	}
}
