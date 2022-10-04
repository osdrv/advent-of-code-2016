package main

import (
	"os"
	"strings"
)

const (
	_ int = iota
	Rect
	RotCol
	RotRow
)

func parseInstr(s string) [3]int {
	var instr [3]int
	if startsWith(s, "rect") {
		instr[0] = Rect
		chs := strings.SplitN(s[5:], "x", 2)
		instr[1], instr[2] = parseInt(chs[0]), parseInt(chs[1])
	} else if startsWith(s, "rotate column") {
		instr[0] = RotCol
		chs := strings.SplitN(s[16:], " by ", 2)
		instr[1], instr[2] = parseInt(chs[0]), parseInt(chs[1])
	} else if startsWith(s, "rotate row") {
		instr[0] = RotRow
		chs := strings.SplitN(s[13:], " by ", 2)
		instr[1], instr[2] = parseInt(chs[0]), parseInt(chs[1])
	} else {
		panic("oops")
	}
	return instr
}

func exec(display [][]int, instrs [][3]int) {
	h, w := sizeNumField(display)
	for _, instr := range instrs {
		switch instr[0] {
		case Rect:
			for i := 0; i < instr[2]; i++ {
				for j := 0; j < instr[1]; j++ {
					display[i][j] = 1
				}
			}
		case RotCol:
			cp := make([]int, h)
			for i := 0; i < h; i++ {
				cp[i] = display[i][instr[1]]
			}
			for i := 0; i < h; i++ {
				display[(i+instr[2])%h][instr[1]] = cp[i]
			}
		case RotRow:
			cp := make([]int, w)
			for j := 0; j < w; j++ {
				cp[j] = display[instr[1]][j]
			}
			for j := 0; j < w; j++ {
				display[instr[1]][(j+instr[2])%w] = cp[j]
			}
		default:
			panic("oopsie doopsie")
		}
		println(printDisplay(display))
	}
}

func printDisplay(display [][]int) string {
	return printNumFieldWithSubs(display, "", map[int]string{
		1: "#",
		0: ".",
	})
}

func countPixels(display [][]int) int {
	cnt := 0
	for i := 0; i < len(display); i++ {
		for j := 0; j < len(display[0]); j++ {
			if display[i][j] > 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	instrs := make([][3]int, 0, len(lines))
	for _, line := range lines {
		instr := parseInstr(line)
		instrs = append(instrs, instr)
	}
	printf("instrs: %+v", instrs)

	//display := makeNumField[int](3, 7)
	display := makeIntField(6, 50)

	exec(display, instrs)

	printf("%d pixels are on", countPixels(display))

}
