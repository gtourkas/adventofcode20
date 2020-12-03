package adventofcode20

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func RepairReport2(expenses []int) int {

	lenExps := len(expenses)
	for i:=0; i<lenExps; i++{
		for j:=i+1; j<lenExps; j++ {
			expI := expenses[i]
			expJ := expenses[j]
			if expI + expJ == 2020 {
				return expI * expJ
			}
		}
	}

	return 0
}

func RepairReport3(expenses []int) int {

	lenExps := len(expenses)
	for i:=0; i<lenExps; i++{
		for j:=i+1; j<lenExps; j++ {
			for k:=j+1; k<lenExps; k++ {
				expI := expenses[i]
				expJ := expenses[j]
				expK := expenses[k]
				if expI + expJ + expK == 2020 {
					return expI * expJ * expK
				}
			}
		}
	}

	return 0
}

type RepairReport func(expenses []int) int

func RunRepairReport(f RepairReport) int {
	file, err := os.Open("day1_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var expenses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expStr := scanner.Text()
		exp, _ := strconv.Atoi(expStr)

		expenses = append(expenses, exp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return f(expenses)
}