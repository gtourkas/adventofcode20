package adventofcode20

import (
	"testing"
)

func TestRunCountValidPassportsExampleWithoutValidation(t *testing.T) {
	act := RunCountValidPassports("day4_example", false)
	exp := 2
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRunCountValidPassportsPuzzleInputWithoutValidation(t *testing.T) {
	act := RunCountValidPassports("day4_input", false)
	exp := 233
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRunCountValidPassportsExampleWithValidation(t *testing.T) {
	act := RunCountValidPassports("day4_example", true)
	exp := 2
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRunCountValidPassportsAllInvalidsInputWithValidation(t *testing.T) {
	act := RunCountValidPassports("day4_allinvalids", true)
	exp := 0
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRunCountValidPassportsAllValidsInputWithValidation(t *testing.T) {
	act := RunCountValidPassports("day4_allvalids", true)
	exp := 4
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRunCountValidPassportsPuzzleInputWithValidation(t *testing.T) {
	act := RunCountValidPassports("day4_input", true)
	exp := 111
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
