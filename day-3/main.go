package main

import (
	"os"
)

func parseTriangle(s string) []int {
	nums := parseInts(s)
	return []int{nums[0], nums[1], nums[2]}
}

func isPossible(tr []int) bool {
	a, b, c := tr[0], tr[1], tr[2]
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	trs := make([][]int, 0, len(lines))
	for _, line := range lines {
		trs = append(trs, parseTriangle(line))
	}

	cnt := 0
	for _, tr := range trs {
		if isPossible(tr) {
			cnt++
		}
	}
	printf("total count: %d", cnt)

	cnt2 := 0
	for y := 0; y < len(trs); y += 3 {
		for x := 0; x < len(trs[y]); x++ {
			tr := []int{trs[y][x], trs[y+1][x], trs[y+2][x]}
			if isPossible(tr) {
				cnt2++
			}
		}
	}
	printf("total count2: %d", cnt2)
}
