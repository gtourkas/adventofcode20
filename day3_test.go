package adventofcode20

import "testing"

func RunCountTreeCases(filename string, moves [][]int) int {
	r := 1
	for i:=0; i<len(moves); i++ {
		move := moves[i]
		f := func(lines []string) int {
			return CountTrees(lines, move[0], move[1])
		}
		r *= RunCountTrees(filename, f)
	}
	return r
}

var part1Case = [][]int{{3,1}}

func TestRunCountTreesExample_Part1Case(t *testing.T) {
	act := RunCountTreeCases("day3_example", part1Case)
	exp := 7
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
func TestRunCountTreesPuzzleInput_Part1Case(t *testing.T) {
	act := RunCountTreeCases("day3_input", part1Case)
	exp := 184
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

var part2Cases = [][]int{
	{1,1},
	part1Case[0],
	{5,1},
	{7,1},
	{1,2},
}

func TestRunCountTreesExample_Part2Cases(t *testing.T) {
	act := RunCountTreeCases("day3_example", part2Cases)
	exp := 336
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
func TestRunCountTreesPuzzleInput_Part2Cases(t *testing.T) {
	act := RunCountTreeCases("day3_input", part2Cases)
	exp := 2431272960
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}
