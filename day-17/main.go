package main

import (
	"crypto/md5"
	"encoding/hex"
)

func getMD5Hash(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

const (
	MAX_X = 3
	MAX_Y = 3
)

func isDoorOpen(b byte) bool {
	return b >= 'b' && b <= 'f'
}

func traverse(input string, exhaustive bool) string {
	type queueItem struct {
		x, y int
		path string
	}

	q := make([]queueItem, 0, 1)
	q = append(q, queueItem{0, 0, ""})

	var maxLen = -ALOT
	var maxPath = ""
	var head queueItem
	for len(q) > 0 {
		head, q = q[0], q[1:]
		if head.x == MAX_X && head.y == MAX_Y {
			if !exhaustive {
				return head.path
			}
			if len(head.path) > maxLen {
				maxPath = head.path
			}
			continue
		}

		h := getMD5Hash(input + head.path)
		if up := isDoorOpen(h[0]); head.y > 0 && up {
			q = append(q, queueItem{head.x, head.y - 1, head.path + "U"})
		}
		if down := isDoorOpen(h[1]); head.y < MAX_Y && down {
			q = append(q, queueItem{head.x, head.y + 1, head.path + "D"})
		}
		if left := isDoorOpen(h[2]); head.x > 0 && left {
			q = append(q, queueItem{head.x - 1, head.y, head.path + "L"})
		}
		if right := isDoorOpen(h[3]); head.x < MAX_X && right {
			q = append(q, queueItem{head.x + 1, head.y, head.path + "R"})
		}
	}

	return maxPath
}

func main() {
	//input := "hijkl"
	//input := "ihgpwlah"
	//input := "kglvqrro"
	//input := "ulqzkmiv"
	input := "gdjjyniy"
	//input := "ulqzkmiv"

	path := traverse(input, false)
	printf("shortest path: %s", path)

	longestPath := traverse(input, true)
	printf("longest path: %d", len(longestPath))
}
