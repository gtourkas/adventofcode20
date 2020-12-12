package adventofcode20

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func HandheldHaltingPartA(instructions []Instruction) int {
	accum, _ := HandheldHalting(instructions)
	return accum
}

func HandheldHaltingPartB(instructions []Instruction) int {
	curNops := 0
	for {
		nops := 0
		cp := make([]Instruction, len(instructions))
		for i, instr := range instructions {
			cpInstr := Instruction{
				Operation: instr.Operation,
				Argument:  instr.Argument,
			}
			if cpInstr.Operation == "jmp" {
				nops++
				if nops == curNops+1 {
					cpInstr.Operation = "nop"
				}
			}
			cp[i] = cpInstr
		}
		accum, infiniteLoop := HandheldHalting(cp)
		if !infiniteLoop {
			return accum
		}
		curNops++
	}

}

func HandheldHalting(instructions []Instruction) (int, bool) {
	executed := make([]bool, len(instructions))
	i := 0
	accum := 0
	for {
		if ok := executed[i]; ok {
			return accum, true
		}
		executed[i] = true
		switch instr := instructions[i]; instr.Operation {
		case "acc":
			accum += instr.Argument
			i++
		case "jmp":
			i += instr.Argument
		default: // nop
			i++
		}
		if i == len(instructions) {
			break
		}
	}
	return accum, false
}

type Instruction struct {
	Operation string
	Argument  int
}

func RunHandheldHalting(fileName string, f func([]Instruction) int) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		cmd := parts[0]
		arg, _ := strconv.Atoi(parts[1])

		instructions = append(instructions, Instruction{
			Operation: cmd,
			Argument:  arg,
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return f(instructions)
}
