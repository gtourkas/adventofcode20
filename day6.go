package adventofcode20

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func CustomCustomsCount(lines []string) int {
	r := 0
	groupAnswers := make([]string, 1)
	groupAnswers[0] = ""

	g := 0
	for _, line := range lines {
		if line == "" {
			g++
			groupAnswers = append(groupAnswers, "")
		}
		for i := range line {
			c := string(line[i])
			groupAnswer := groupAnswers[g]
			if !strings.Contains(groupAnswer, c) {
				groupAnswer += c
				groupAnswers[g] = groupAnswer
				r++
			}
		}
	}

	return r
}

func CustomCustomsAllAnswered(lines []string) int {
	r := 0

	groupPeople := 0
	groupAnswers := make(map[uint8]int)

	g := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {

			for c := range groupAnswers {
				if groupAnswers[c] == groupPeople {
					r++
				}
			}

			groupPeople = 0
			groupAnswers = make(map[uint8]int)
			g++
			continue
		}
		for j := range line {
			c := line[j]
			if _, ok := groupAnswers[c]; ok {
				groupAnswers[c]++
			} else {
				groupAnswers[c] = 1
			}
		}
		groupPeople++
	}
	for c := range groupAnswers {
		if groupAnswers[c] == groupPeople {
			r++
		}
	}

	return r
}

func RunCustomCustoms(fileName string, f func([]string) int) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return f(lines)
}
