package main

import (
	"os"
	"strings"
)

func dist(p1, p2 Point2) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

/*
    rotRight:
	b -a
	0 1
	1 0
	0 -1
	-1 0

	rotLeft:
	-b a
	0 1
	-1 0
	0 -1
	1 0
*/

func rotLeft(dir [2]int) [2]int {
	return [2]int{-dir[1], dir[0]}
}

func rotRight(dir [2]int) [2]int {
	return [2]int{dir[1], -dir[0]}
}

func traversePath(pos Point2, steps []string) Point2 {
	dir := [2]int{0, 1}
	visited := make(map[Point2]bool)
	visited[pos] = true
	found := false
	for _, step := range steps {
		if step[0] == 'R' {
			dir = rotRight(dir)
		} else if step[0] == 'L' {
			dir = rotLeft(dir)
		} else {
			panic("wtf")
		}
		dd := parseInt(step[1:])
		dx := dd * dir[0]
		dy := dd * dir[1]
		if !found {
			for y := 0; y <= dy; y++ {
				for x := 0; x <= dx; x++ {
					if x == 0 && y == 0 {
						continue
					}
					pp := Point2{pos.x + x, pos.y + y}
					if _, ok := visited[pp]; ok {
						printf("first location visited twice: %+v (dist: %d)", pp, dist(Point2{0, 0}, pp))
						found = true
					}
					visited[pp] = true
				}
			}
		}
		pos.x += dx
		pos.y += dy
	}

	return pos
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	for _, line := range lines {
		steps := strings.Split(line, ", ")
		pos := traversePath(Point2{0, 0}, steps)
		printf("pos: %+v", pos)
		printf("dist: %d", dist(Point2{0, 0}, pos))
	}
}
