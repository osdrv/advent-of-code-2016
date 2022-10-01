package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	doorId := "cxdnnyjw"
	//doorId := "abc"
	var b bytes.Buffer

	cnt := 0
	i := 0
	for cnt < 8 {
		d := doorId + strconv.Itoa(i)
		s := fmt.Sprintf("%x", md5.Sum([]byte(d)))
		if s[:5] == "00000" {
			b.WriteByte(s[5])
			cnt++
		}
		i++
	}
	printf(b.String())

	pwd := make([]byte, 8)
	tr := 0b00000000
	trl := 0b11111111
	i2 := 0
	for tr != trl {
		d := doorId + strconv.Itoa(i2)
		s := fmt.Sprintf("%x", md5.Sum([]byte(d)))
		if s[:5] == "00000" {
			printf("hash: %s", s)
			pos := s[5]
			char := s[6]
			if pos >= '0' && pos <= '7' {
				if tr&(1<<int(pos-'0')) == 0 {
					pwd[int(pos-'0')] = char
					tr |= (1 << int(pos-'0'))
					printf("%c goes to %c, cnt: %08b", char, pos, tr)
				}
			}
		}
		i2++
	}
	printf("password2: %s", string(pwd))
}
