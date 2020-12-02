package adventofcode20

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type PasswordFileEntry struct {
	Min int
	Max int
	Char uint8
	Password string
}

type CountPasswords func(entries []PasswordFileEntry) int

func CountValidPasswordsTheOldWay(entries []PasswordFileEntry) int {
	r := 0
	for _,e := range entries {
		c := 0
		for i := range e.Password {
			if e.Password[i] == e.Char {
				c++
			}
		}
		if c >= e.Min && c <= e.Max {
			r ++
		}
	}
	return r
}

func CountValidPasswordsTheNewWay(entries []PasswordFileEntry) int {
	r := 0
	for _,e := range entries {
		lenPassword := len(e.Password)
		if e.Min > lenPassword || e.Max > lenPassword {
			continue
		}
		chAtMin := e.Password[e.Min-1]
		chAtMax := e.Password[e.Max-1]
		if (chAtMin == e.Char && chAtMax != e.Char) ||
			(chAtMin != e.Char && chAtMax == e.Char){
			r ++
		}
	}
	return r
}

func RunCountValidPasswords(f CountPasswords) int {
	file, err := os.Open("day2_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`^(?P<min>\d+)-(?P<max>\d+)\s(?P<char>[a-z]):\s(?P<password>[a-z]+)$`)

	var entries []PasswordFileEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, match := range re.FindAllStringSubmatch(line, -1) {

			min, _ := strconv.Atoi(match[1])
			max, _ := strconv.Atoi(match[2])
			char := match[3]
			password := match[4]
			entry := PasswordFileEntry{
				Min:      min,
				Max:      max,
				Char:     char[0],
				Password: password,
			}

			entries = append(entries, entry)
		}
	}

	r := f(entries)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return r
}