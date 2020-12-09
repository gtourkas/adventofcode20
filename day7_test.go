package adventofcode20

import "testing"

func TestHandyHaversacksPartAWithExample(t *testing.T) {
	act := RunHandyHaversacks("day7_example", HandyHaversacksPartA)
	exp := 4
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandyHaversacksPartAWithPuzzleInput(t *testing.T) {
	act := RunHandyHaversacks("day7_input", HandyHaversacksPartA)
	exp := 121
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandyHaversacksPartBWithExample(t *testing.T) {
	act := RunHandyHaversacks("day7_example", HandyHaversacksPartB)
	exp := 32
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandyHaversacksPartBWithExample2(t *testing.T) {
	act := RunHandyHaversacks("day7_example2", HandyHaversacksPartB)
	exp := 126
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandyHaversacksPartBWithPuzzleInput(t *testing.T) {
	act := RunHandyHaversacks("day7_input", HandyHaversacksPartB)
	exp := 3805
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
