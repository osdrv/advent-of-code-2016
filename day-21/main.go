package main

import (
	"os"
)

func parseSwapPositions(s string) (int, int) {
	// swap position X with position Y
	return int(s[14] - '0'), int(s[30] - '0')
}

func parseSwapLetters(s string) (byte, byte) {
	// swap letter X with letter Y
	return s[12], s[26]
}

func parseRotate(s string) (string, int) {
	// rotate left/right X steps
	ptr := 0
	_, ptr = readStr(s, ptr, "rotate ")
	var dir string
	dir, ptr = readWord(s, ptr)
	ptr = eatWhitespace(s, ptr)
	var off int
	off, ptr = readInt(s, ptr)
	return dir, off
}

func parseRotateAround(s string) byte {
	// rotate based on position of letter X
	return s[35]
}

func parseReverse(s string) (int, int) {
	// reverse positions X through Y
	return int(s[18] - '0'), int(s[28] - '0')
}

func parseMovePos(s string) (int, int) {
	// move position X to position Y
	return int(s[14] - '0'), int(s[28] - '0')
}

func swapPos(bb []byte, p1, p2 int) []byte {
	bbcp := make([]byte, len(bb))
	copy(bbcp, bb)
	bbcp[p1], bbcp[p2] = bbcp[p2], bbcp[p1]
	return bbcp
}

func findChar(bb []byte, ch byte) int {
	for i := 0; i < len(bb); i++ {
		if bb[i] == ch {
			return i
		}
	}
	return -1
}

func swapLetters(bb []byte, c1, c2 byte) []byte {
	return swapPos(bb, findChar(bb, c1), findChar(bb, c2))
}

func rotate(bb []byte, dir string, off int) []byte {
	off %= len(bb)
	if dir == "right" {
		off = len(bb) - off
	}
	bbcp := make([]byte, len(bb))
	copy(bbcp, bb[off:])
	copy(bbcp[len(bb[off:]):], bb[:off])
	return bbcp
}

func rotateAround(bb []byte, ch byte) []byte {
	p := findChar(bb, ch)
	off := 1 + p
	if p >= 4 {
		off++
	}
	return rotate(bb, "right", off)
}

/*

  abcdefgh
  b
  >> 1 + 1
  ghabcdef
  b
  << 3 - 1

  abcdefgh
  f
  >> 1 + 5 + 1 = 6
  cdefghab

*/

//func unrotateAround(bb []byte, ch byte) []byte {
//	p := findChar(bb, ch)
//	p--
//
//}

func reverse(bb []byte, p1, p2 int) []byte {
	bbcp := make([]byte, len(bb))
	copy(bbcp, bb)
	if p1 > p2 {
		p1, p2 = p2, p1
	}
	for p1 < p2 {
		bbcp[p1], bbcp[p2] = bbcp[p2], bbcp[p1]
		p1++
		p2--
	}
	return bbcp
}

func movePos(bb []byte, p1, p2 int) []byte {
	bbcp := make([]byte, len(bb))
	bbcp[p2] = bb[p1]
	ptr1, ptr2 := 0, 0
	for ptr1 < len(bb) {
		if ptr1 == p1 {
			ptr1++
			continue
		}
		if ptr2 == p2 {
			ptr2++
			continue
		}
		bbcp[ptr2] = bb[ptr1]
		ptr1++
		ptr2++
	}
	return bbcp
}

func scramble(s string, instrs []string) string {
	bb := []byte(s)

	for _, instr := range instrs {
		debugf("applying instr: %s", instr)
		debugf("input string: %s", string(bb))
		if matchStr(instr, 0, "swap position") {
			p1, p2 := parseSwapPositions(instr)
			bb = swapPos(bb, p1, p2)
		} else if matchStr(instr, 0, "swap letter") {
			c1, c2 := parseSwapLetters(instr)
			bb = swapLetters(bb, c1, c2)
		} else if matchStr(instr, 0, "rotate left") || matchStr(instr, 0, "rotate right") {
			dir, off := parseRotate(instr)
			bb = rotate(bb, dir, off)
		} else if matchStr(instr, 0, "rotate based on position of letter") {
			ch := parseRotateAround(instr)
			bb = rotateAround(bb, ch)
		} else if matchStr(instr, 0, "reverse positions") {
			p1, p2 := parseReverse(instr)
			bb = reverse(bb, p1, p2)
		} else if matchStr(instr, 0, "move position") {
			p1, p2 := parseMovePos(instr)
			bb = movePos(bb, p1, p2)
		}
		debugf("resulting string: %s", string(bb))
	}

	return string(bb)
}

func genAll(n int) []string {
	var recurse func(int, []bool) []string
	recurse = func(j int, taken []bool) []string {
		if j == n {
			return []string{""}
		}
		res := make([]string, 0, 1)
		for i := 0; i < len(taken); i++ {
			if taken[i] {
				continue
			}
			taken[i] = true
			for _, opt := range recurse(j+1, taken) {
				res = append(res, string(append([]byte{byte(i) + 'a'}, []byte(opt)...)))
			}
			taken[i] = false
		}
		return res
	}

	return recurse(0, make([]bool, n))
}

func unscramble(s string, lines []string) string {
	candidates := genAll(len(s))
	for _, c := range candidates {
		debugf("inspecting candidate: %s", c)
		if scramble(c, lines) == s {
			return c
		}
	}
	panic("nothing found")
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	//res := scramble("abcde", lines)
	res := scramble("abcdefgh", lines)
	printf("scrambled password: %s", res)

	res2 := unscramble("fbgdceah", lines)
	printf("unscrambled password: %s", res2)
}
