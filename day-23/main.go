package main

import (
	"os"
	"strings"
)

func parseInstr(s string) []string {
	return strings.Split(s, " ")
}

func getVal(regs map[string]int, arg string) int {
	if looksLikeAddr(arg) {
		return regs[arg]
	}
	return parseInt(arg)
}

func looksLikeAddr(arg string) bool {
	return len(arg) == 1 && arg[0] >= 'a' && arg[0] <= 'z'
}

func interpret(instrs [][]string, regs map[string]int) {
	pc := 0
	for pc < len(instrs) {
		i := instrs[pc]
		debugf("interpret %+v", i)
		debugf("regs: %+v", regs)
		switch i[0] {
		case "cpy":
			if looksLikeAddr(i[2]) {
				regs[i[2]] = getVal(regs, i[1])
			}
			pc++
		case "inc":
			if looksLikeAddr(i[1]) {
				regs[i[1]]++
			}
			pc++
		case "dec":
			if looksLikeAddr(i[1]) {
				regs[i[1]]--
			}
			pc++
		case "jnz":
			if x := getVal(regs, i[1]); x != 0 {
				y := getVal(regs, i[2])
				pc += y
			} else {
				pc++
			}
		case "tgl":
			off := getVal(regs, i[1])
			if pc+off < 0 || pc+off >= len(instrs) {
				pc++
				continue
			}
			ii := instrs[pc+off]
			if len(ii) == 2 {
				// single-argument instr
				if ii[0] == "inc" {
					ii[0] = "dec"
				} else {
					ii[0] = "inc"
				}
			} else if len(ii) == 3 {
				if ii[0] == "jnz" {
					ii[0] = "cpy"
				} else {
					ii[0] = "jnz"
				}
			}
			pc++
		default:
			panic("wtf")
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
	regs["a"] = 12

	interpret(instrs, regs)

	printf("regs: %+v", regs)
}
