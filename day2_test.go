package adventofcode20

import "testing"

func TestRunCountValidPasswordsTheOldWayWithExample(t *testing.T) {
	act := CountValidPasswordsTheOldWay([]PasswordFileEntry {
		{Min: 1, Max: 3, Char: 'a', Password: "abcde"},
		{Min: 1, Max: 3, Char: 'b', Password: "cdefg"},
		{Min: 2, Max: 9, Char: 'c', Password: "ccccccccc"},
	})
	exp := 2
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}


func TestRunCountValidPasswordsTheOldWayWithPuzzleInput(t *testing.T) {
	act := RunCountValidPasswords(CountValidPasswordsTheOldWay)
	exp := 586
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRunCountValidPasswordsTheNewWayWithExample(t *testing.T) {
	act := CountValidPasswordsTheNewWay([]PasswordFileEntry {
		{Min: 1, Max: 3, Char: 'a', Password: "abcde"},
		{Min: 1, Max: 3, Char: 'b', Password: "cdefg"},
		{Min: 2, Max: 9, Char: 'c', Password: "ccccccccc"},
	})
	exp := 1
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}


func TestRunCountValidPasswordsTheNewWayWithPuzzleInput(t *testing.T) {
	act := RunCountValidPasswords(CountValidPasswordsTheNewWay)
	exp := 352
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
