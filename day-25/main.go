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

func interpret(instrs [][]string, regs map[string]int) bool {
	prev := -1
	cnt := 0
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
		case "out":
			v := getVal(regs, i[1])
			if v == prev {
				return false
			}
			prev = v
			cnt++
			if cnt >= 10 {
				return true
			}
		}
	}
	return true
}

func arrEql(a1, a2 []int) bool {
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func routine(a int) bool {
	prev := -1
	cnt := 0
	res := make([]int, 0, 10)
	var b, c, d int
	d = a
	c = 11
L2:
	b = 231
L1:
	d++
	b--
	if b != 0 {
		goto L1
	}
	c--
	if c != 0 {
		goto L2
	}
L11:
	a = d
L10:
	b = a
	a = 0
L6:
	c = 2
L5:
	if b != 0 {
		goto L3
	}
	goto L4
L3:
	b--
	c--
	if c != 0 {
		goto L5
	}
	a++
	goto L6
L4:
	b = 2
L9:
	if c != 0 {
		goto L7
	}
	goto L8
L7:
	b--
	c--
	goto L9
L8:
	if b == prev {
		return false
	}
	res = append(res, b)
	prev = b
	cnt++
	if cnt >= 10 {
		printf("res: %+v", res)
		return true
	}
	if a != 0 {
		goto L10
	}
	goto L11
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

	a := 0
	for {
		//regs := make(map[string]int)
		//regs["a"] = a
		//if !interpret(instrs, regs) {
		if !routine(a) {
			a++
			if a%1000 == 0 {
				printf("checkpoint at a: %d", a)
			}
			continue
		}
		printf("a: %d", a)
		break
	}
}
