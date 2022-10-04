package main

import (
	"os"
	"strings"
)

func parseIP(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '[' || r == ']'
	})
}

func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if (s[i] == s[i+3]) && (s[i+1] == s[i+2]) && (s[i] != s[i+1]) {
			return true
		}
	}
	return false
}

func isSupportsTLS(ip []string) bool {
	for i := 1; i < len(ip); i = i + 2 {
		if containsABBA(ip[i]) {
			debugf("%s should not contain ABBA", ip[i])
			return false
		}
	}
	for i := 0; i < len(ip); i += 2 {
		if containsABBA(ip[i]) {
			return true
		}
	}
	return false
}

func isSupportsSSL(ip []string) bool {
	debugf("ip: %+v", ip)
	babs := make(map[string]bool)
	for i := 0; i < len(ip); i += 2 {
		s := ip[i]
		for j := 0; j < len(s)-2; j++ {
			if s[j] == s[j+2] && s[j] != s[j+1] {
				bab := string([]byte{s[j+1], s[j], s[j+1]})
				babs[bab] = true
			}
		}
	}
	debugf("babs: %+v", babs)
	for i := 1; i < len(ip); i += 2 {
		s := ip[i]
		for j := 0; j < len(s)-2; j++ {
			if _, ok := babs[s[j:j+3]]; ok {
				return true
			}
		}
	}
	return false
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	cntTLS := 0
	cntSSL := 0
	for _, line := range lines {
		ip := parseIP(line)
		supTLS := isSupportsTLS(ip)
		supSSL := isSupportsSSL(ip)
		printf("s: %s, sup TLS: %t, sup SSL: %t", line, supTLS, supSSL)
		if supTLS {
			cntTLS++
		}
		if supSSL {
			cntSSL++
		}
	}

	printf("count TLS: %d", cntTLS)
	printf("count SSL: %d", cntSSL)
}
