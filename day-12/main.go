package main

import (
	"os"
	"strings"
)

func parseInstr(s string) []string {
	return strings.Split(s, " ")
}

func getVal(regs map[string]int, arg string) int {
	if len(arg) == 1 && arg[0] >= 'a' && arg[0] <= 'z' {
		return regs[arg]
	}
	return parseInt(arg)
}

func interpret(instrs [][]string, regs map[string]int) {
	pc := 0
	for pc < len(instrs) {
		i := instrs[pc]
		debugf("interpret %+v", i)
		debugf("regs: %+v", regs)
		switch i[0] {
		case "cpy":
			regs[i[2]] = getVal(regs, i[1])
			pc++
		case "inc":
			regs[i[1]]++
			pc++
		case "dec":
			regs[i[1]]--
			pc++
		case "jnz":
			if x := getVal(regs, i[1]); x != 0 {
				y := getVal(regs, i[2])
				pc += y
			} else {
				pc++
			}
		}
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	instrs := make([][]string, 0, len(lines))

	for _, line := range lines {
		instrs = append(instrs, parseInstr(line))
	}

	regs := make(map[string]int)

	regs["c"] = 1

	interpret(instrs, regs)

	printf("regs: %+v", regs)
}
