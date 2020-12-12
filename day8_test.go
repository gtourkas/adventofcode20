package adventofcode20

import "testing"

func TestHandheldHaltingPartAWithExample(t *testing.T) {
	act := RunHandheldHalting("day8_example", HandheldHaltingPartA)
	exp := 5
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandheldHaltingPartAWithExample2(t *testing.T) {
	act := RunHandheldHalting("day8_example2", HandheldHaltingPartA)
	exp := 8
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandheldHaltingPartAWithPuzzleInput(t *testing.T) {
	act := RunHandheldHalting("day8_input", HandheldHaltingPartA)
	exp := 2034
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandheldHaltingPartBWithExample(t *testing.T) {
	act := RunHandheldHalting("day8_example", HandheldHaltingPartB)
	exp := 8
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestHandheldHaltingPartBWithPuzzleInput(t *testing.T) {
	act := RunHandheldHalting("day8_input", HandheldHaltingPartB)
	exp := 672
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
