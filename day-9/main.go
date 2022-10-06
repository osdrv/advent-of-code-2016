package main

import (
	"bytes"
	"os"
)

func parseQuant(s string, ptr int) ([2]int, int) {
	var sub, rep int
	ptr = consume(s, ptr, '(')
	sub, ptr = readInt(s, ptr)
	ptr = consume(s, ptr, 'x')
	rep, ptr = readInt(s, ptr)
	ptr = consume(s, ptr, ')')
	return [2]int{sub, rep}, ptr
}

func expand2(s string) int {
	cnt := 0

	ptr := 0
	for ptr < len(s) {
		for ptr < len(s) && isAlpha(s[ptr]) {
			cnt++
			ptr++
		}
		if match(s, ptr, '(') {
			var q [2]int
			q, ptr = parseQuant(s, ptr)
			explen := expand2(s[ptr : ptr+q[0]])
			cnt += q[1] * explen
			ptr += q[0]
		}
	}

	return cnt
}

func expand(s string) string {
	var b bytes.Buffer
	ptr := 0
	for ptr < len(s) {
		if match(s, ptr, '(') {
			var q [2]int
			q, ptr = parseQuant(s, ptr)
			for i := 0; i < q[1]; i++ {
				b.WriteString(s[ptr : ptr+q[0]])
			}
			ptr += q[0]
		}
		if ptr >= len(s) {
			break
		}
		from := ptr
		for ptr < len(s) && isAlpha(s[ptr]) {
			ptr++
		}
		b.WriteString(s[from:ptr])
	}

	return b.String()
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	for _, line := range lines {
		exp := expand(line)
		printf("line: %s", line)
		printf("expanded: %s, len: %d", exp, len(exp))
		exp2 := expand2(line)
		printf("expanded2 len: %d", exp2)
	}
}
