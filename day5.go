package adventofcode20

import (
	"bufio"
	"log"
	"os"
)

func BinaryBoarding(seat string) int {

	lowerRow := 0
	upperRow := 127
	leftCol := 0
	rightCol := 7
	for i := range seat {
		c := seat[i]
		if c == 'F' || c == 'B' {
			h := (upperRow - lowerRow) / 2
			if c == 'F' {
				upperRow -= h + 1
			}
			if c == 'B' {
				lowerRow += h + 1
			}
		} else {
			h := (rightCol - leftCol) / 2
			if c == 'L' {
				rightCol -= h + 1
			}
			if c == 'R' {
				leftCol += h + 1
			}
		}
	}

	return upperRow*8 + rightCol
}

func RunBinaryBoarding(fileName string, f func(line string)) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
