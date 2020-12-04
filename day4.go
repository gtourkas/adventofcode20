package adventofcode20

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func CreatePassportEntry(lines []string) map[string]string {
	r := make(map[string]string)
	for _, l := range lines {
		entries := strings.Split(l, " ")
		for _, e := range entries {
			parts := strings.Split(e, ":")
			r[parts[0]] = parts[1]
		}
	}
	return r
}

type FieldCfg struct {
	Mandatory        bool
	ValidationRegExp *regexp.Regexp
}

var PassportFields = map[string]FieldCfg{
	"byr": {
		true,
		regexp.MustCompile("^(19[2-9][0-9]|200[0-2])$"),
	},
	"iyr": {
		true,
		regexp.MustCompile("^(201[0-9]|2020)$"),
	},
	"eyr": {
		true,
		regexp.MustCompile("^(202[0-9]|2030)$"),
	},
	"hgt": {
		true,
		regexp.MustCompile("^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$"),
	},
	"hcl": {
		true,
		regexp.MustCompile("^#([0-9]|[a-f]){6}$"),
	},
	"ecl": {
		true,
		regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$"),
	},
	"pid": {
		true,
		regexp.MustCompile("^[0-9]{9}$"),
	},
	"cid": {
		false,
		nil,
	},
}

func CountValidPassports(entries []map[string]string, validate bool) int {
	r := 0
	for _, e := range entries {
		valid := true
		for f, cfg := range PassportFields {
			if value, ok := e[f]; (!ok && cfg.Mandatory) ||
				(validate && cfg.ValidationRegExp != nil && !cfg.ValidationRegExp.Match([]byte(value))) {
				valid = false
				break
			}
		}
		if valid {
			r++
		}
	}

	return r
}

func RunCountValidPassports(fileName string, validate bool) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var entries []map[string]string
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			entry := CreatePassportEntry(lines)
			entries = append(entries, entry)
			lines = []string{}
			continue
		}

		lines = append(lines, line)
	}

	if len(lines) > 0 {
		entry := CreatePassportEntry(lines)
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return CountValidPassports(entries, validate)
}
