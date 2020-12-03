package adventofcode20

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func CountTrees(lines []string, right int, down int) int {
	trees := 0

	r := 0
	lenLine := len(lines[0])
	extends := 0
	for i:=0; i<len(lines); i+=down {
		line := lines[i]

		if extends > 0 {
			line = strings.Repeat(line, extends+1)
			lenLine = len(line)
		}

		// extend further the line if you're too right
		if r >= lenLine {
			extends ++
			line = strings.Repeat(line, extends+1)
			lenLine = len(line)
		}
		c := line[r]

		if c == '#' {
			trees ++
		}

		r += right
	}

	return trees
}

type CountTreesClosure func(lines []string) int

func RunCountTrees(fileName string, f CountTreesClosure) int {
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
