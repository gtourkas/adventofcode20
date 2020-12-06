package adventofcode20

import (
	"reflect"
	"testing"
)

func TestBinaryBoardingWithExamples(t *testing.T) {
	var tests = []struct {
		name     string
		seat     string
		expected int
	}{
		{"Case 1", "FBFBBFFRLR", 357},
		{"Case 2", "BFFFBBFRRR", 567},
		{"Case 3", "FFFBBBFRRR", 119},
		{"Case 4", "BBFFBBFRLL", 820},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := BinaryBoarding(tt.seat)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("failed test: %s // expected: %d // actual: %d", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestBinaryBoardingWithPuzzleInputToFindMaxSeatID(t *testing.T) {
	act := 0
	f := func(line string) {
		r := BinaryBoarding(line)
		if r > act {
			act = r
		}
	}
	RunBinaryBoarding("day5_input", f)
	exp := 816
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestBinaryBoardingWithPuzzleInputToFindMySeatID(t *testing.T) {
	act := 0
	seats := make([]bool, 1024)
	f := func(line string) {
		r := BinaryBoarding(line)
		seats[r-1] = true
	}
	RunBinaryBoarding("day5_input", f)
	var prev, cur, next bool
	for i := range seats {
		next = seats[i]
		if prev && !cur && next {
			act = i // meaning cur
			break
		}
		prev = cur
		cur = next
	}
	exp := 539
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
