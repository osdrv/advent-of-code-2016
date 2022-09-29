package main

import (
	"bytes"
	"os"
	"strconv"
)

var (
	STEPS = map[rune][2]int{
		'U': {0, -1},
		'D': {0, 1},
		'L': {-1, 0},
		'R': {1, 0},
	}
)

const (
	A int = 10 + iota
	B
	C
	D
)

var (
	NUMPAD_3x3 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	NUMPAD_5x5 = [][]int{
		{0, 0, 1, 0, 0},
		{0, 2, 3, 4, 0},
		{5, 6, 7, 8, 9},
		{0, A, B, C, 0},
		{0, 0, D, 0, 0},
	}
)

func computeCode(numpad [][]int, pos Point2, ss []string) string {
	var b bytes.Buffer
	w, h := sizeNumField(numpad)
	for _, s := range ss {
		for _, r := range s {
			nx, ny := pos.x+STEPS[r][0], pos.y+STEPS[r][1]
			nx = max(0, min(nx, w-1))
			ny = max(0, min(ny, h-1))
			if numpad[ny][nx] == 0 {
				continue
			}
			pos.x = nx
			pos.y = ny
		}
		var numstr string
		if num := numpad[pos.y][pos.x]; num < A {
			numstr = strconv.Itoa(num)
		} else {
			numstr = string([]byte{'A' + byte(num-A)})
		}
		b.WriteString(numstr)
	}
	return b.String()
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	code := computeCode(NUMPAD_3x3, Point2{1, 1}, lines)

	printf("code: %s", code)

	code2 := computeCode(NUMPAD_5x5, Point2{0, 2}, lines)
	printf("code2: %s", code2)
}
