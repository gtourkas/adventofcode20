package adventofcode20

import "testing"

func TestCustomCustomsCountWithExampleInput(t *testing.T) {
	act := RunCustomCustoms("day6_example", CustomCustomsCount)
	exp := 11
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestCustomCustomsCountWithPuzzleInput(t *testing.T) {
	act := RunCustomCustoms("day6_input", CustomCustomsCount)
	exp := 6170
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestCustomCustomsAllAnsweredWithExampleInput(t *testing.T) {
	act := RunCustomCustoms("day6_example", CustomCustomsAllAnswered)
	exp := 6
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestCustomCustomsAllAnsweredWithPuzzleInput(t *testing.T) {
	act := RunCustomCustoms("day6_input", CustomCustomsAllAnswered)
	exp := 2947
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
