package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func getMd5(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

var (
	memo = make(map[string]string)
)

func getMd5Stretch(s string) string {
	if v, ok := memo[s]; ok {
		return v
	}
	hs := s
	for i := 0; i < 2017; i++ {
		h := md5.Sum([]byte(hs))
		hs = hex.EncodeToString(h[:])
	}
	memo[s] = hs
	return hs
}

func NewHashStream(salt string, hh func(string) string) func() int {
	index := 0

	getNextTrpl := func() (int, byte) {
		for {
			h := hh(salt + strconv.Itoa(index))
			curix := index
			index++

			i := 0
			for i < len(h) {
				j := i + 1
				for j < len(h) && h[j] == h[i] {
					j++
					if j-i >= 3 {
						debugf("proposing candidate hash: %s", h)
						return curix, h[i]
					}
				}
				i = j
			}
		}
	}

	contains5 := func(ix int, ch byte) bool {
		h := hh(salt + strconv.Itoa(ix))
		i := 0
		for i < len(h) {
			for i < len(h) && h[i] != ch {
				i++
			}
			if i >= len(h) {
				break
			}
			j := i + 1
			for j < len(h) && h[j] == h[i] {
				j++
				if j-i >= 5 {
					return true
				}
			}
			i = j
		}
		return false
	}

	return func() int {
		for {
			trix, ch := getNextTrpl()
			debugf("triple candidate: %d (contains %c)", trix, ch)
			for ix := trix + 1; ix <= trix+1000; ix++ {
				if contains5(ix, ch) {
					debugf("contains 5 at index: %d", ix)
					return trix
				}
			}
			debugf("discarded")
		}
	}
}

func main() {
	//SALT := "abc"
	SALT := "qzyelonm"

	res := make([]int, 0, 64)

	hs := NewHashStream(SALT, getMd5Stretch)

	for len(res) < 64 {
		res = append(res, hs())
	}

	printf("The last index is %d", res[len(res)-1])
}
