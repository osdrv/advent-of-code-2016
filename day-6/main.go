package main

import (
	"bytes"
	"os"
)

func computeLeastCommonChars(lines []string) string {
	var b bytes.Buffer
	for col := 0; col < len(lines[0]); col++ {
		var cnts [26]int
		for row := 0; row < len(lines); row++ {
			cnts[int(lines[row][col]-'a')]++
		}
		mincnt := ALOT
		var minch byte
		for i := 0; i < len(cnts); i++ {
			if cnts[i] > 0 && cnts[i] < mincnt {
				mincnt = cnts[i]
				minch = 'a' + byte(i)
			}
		}
		b.WriteByte(minch)
	}
	return b.String()
}

func computeMostCommonChars(lines []string) string {
	var b bytes.Buffer
	for col := 0; col < len(lines[0]); col++ {
		var cnts [26]int
		for row := 0; row < len(lines); row++ {
			cnts[int(lines[row][col]-'a')]++
		}
		maxcnt := 0
		var maxch byte
		for i := 0; i < len(cnts); i++ {
			if cnts[i] > maxcnt {
				maxcnt = cnts[i]
				maxch = 'a' + byte(i)
			}
		}
		b.WriteByte(maxch)
	}
	return b.String()
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	mostCommonChars := computeMostCommonChars(lines)

	printf("most common chars: %s", mostCommonChars)

	leastCommonChars := computeLeastCommonChars(lines)

	printf("least common chars: %s", leastCommonChars)
}
