package main

import (
	"bytes"
	"os"
	"sort"
)

type Room struct {
	name     string
	sectorId int
	checksum string
}

func NewRoom(s string) *Room {
	var ptr int
	var name string
	var sectorId int
	var checksum string
	name, ptr = readStr(s, ptr)
	sectorId, ptr = readNum(s, ptr)
	ptr = consume(s, ptr, '[')
	checksum, ptr = readStr(s, ptr)
	ptr = consume(s, ptr, ']')
	return &Room{
		name:     name,
		sectorId: sectorId,
		checksum: checksum,
	}
}

func readNum(s string, ptr int) (int, int) {
	from := ptr
	for ptr < len(s) && isNum(s[ptr]) {
		ptr++
	}
	return parseInt(s[from:ptr]), ptr
}

func consume(s string, ptr int, b byte) int {
	assert(s[ptr] == b, "expect symbol")
	return ptr + 1
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func isDash(b byte) bool {
	return b == '-'
}

func readStr(s string, ptr int) (string, int) {
	var b bytes.Buffer
	for ptr < len(s) && isAlpha(s[ptr]) || isDash(s[ptr]) {
		if !isDash(s[ptr]) {
			b.WriteByte(s[ptr])
		}
		ptr++
	}
	return b.String(), ptr
}

func (r *Room) IsReal() bool {
	var cnts [26]int
	for _, ch := range r.name {
		cnts[ch-'a']++
	}
	chs := make([]byte, 0, 26)
	for i, cnt := range cnts {
		if cnt > 0 {
			chs = append(chs, 'a'+byte(i))
		}
	}
	sort.Slice(chs, func(i, j int) bool {
		cnt1, cnt2 := cnts[chs[i]-'a'], cnts[chs[j]-'a']
		if cnt1 == cnt2 {
			return chs[i] < chs[j]
		}
		return cnt1 > cnt2
	})
	debugf("cnts: %+v", cnts)
	checksum := string(chs[:5])
	res := checksum == r.checksum
	debugf("name: %s, checksum: %s(%s), res: %t", r.name, checksum, r.checksum, res)
	return res
}

func decypher(r *Room) string {
	var b bytes.Buffer
	for i := 0; i < len(r.name); i++ {
		b.WriteByte('a' + (((r.name[i] - 'a') + byte(r.sectorId%26)) % 26))
	}
	return b.String()
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	rooms := make([]*Room, 0, len(lines))
	for _, line := range lines {
		rooms = append(rooms, NewRoom(line))
	}

	sum := 0
	for _, room := range rooms {
		if room.IsReal() {
			dd := decypher(room)
			if dd == "northpoleobjectstorage" {
				printf("decoded: %s", decypher(room))
				printf("room: %+v", room)
			}
			sum += room.sectorId
		}
	}

	printf("sum: %d", sum)
}
