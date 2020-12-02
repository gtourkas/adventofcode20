package adventofcode20

import (
	"testing"
)

func TestReportReport2WithExample(t *testing.T) {
	act := RepairReport2([]int{1721,
		979,
		366,
		299,
		675,
		1456})
	exp := 514579
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRepairReport2WithPuzzleInput(t *testing.T) {
	act := RunRepairReport(RepairReport2)
	exp := 1019571
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestReportReport3WithExample(t *testing.T) {
	act := RepairReport3([]int{1721,
		979,
		366,
		299,
		675,
		1456})
	exp := 241861950
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}

func TestRepairReport3WithPuzzleInput(t *testing.T) {
	act := RunRepairReport(RepairReport3)
	exp := 100655544
	if act != exp {
		t.Errorf("fail: expected %d / actual %d", exp, act)
	}
}