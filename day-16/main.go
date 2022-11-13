package main

import "bytes"

func dragonCurve(s string) string {
	var b bytes.Buffer
	b.WriteString(s)
	b.WriteByte('0')
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			b.WriteByte('1')
		} else {
			b.WriteByte('0')
		}
	}
	return b.String()
}

func checksum(s string) string {
	var b bytes.Buffer
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] {
			b.WriteByte('1')
		} else {
			b.WriteByte('0')
		}
		i += 2
	}
	return b.String()
}

func main() {
	//seed, size := "10000", 20
	//seed, size := "00111101111101000", 272
	seed, size := "00111101111101000", 35651584

	rndData := seed
	for len(rndData) < size {
		rndData = dragonCurve(rndData)
	}
	rndData = rndData[:size]

	csum := rndData
	for {
		csum = checksum(csum)
		if len(csum)%2 == 1 {
			break
		}
	}

	printf("checksum: %s", csum)
}
