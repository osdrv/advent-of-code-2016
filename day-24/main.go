package main

import (
	"os"
)

func makeField(lines []string) ([][]int, map[int]Point2) {
	field := makeIntField(len(lines), len(lines[0]))
	markers := make(map[int]Point2)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			ch := lines[y][x]
			switch {
			case ch == '#':
				field[y][x] = 0
			case ch == '.':
				field[y][x] = 1
			case ch >= '0' && ch <= '9':
				field[y][x] = 1
				markers[int(ch-'0')] = Point2{x, y}
			}
		}
	}
	return field, markers
}

func findShortestPath(field [][]int, start, target Point2) int {
	height, width := sizeNumField(field)
	steps := make(map[Point2]int)
	q := make([]Point2, 0, 1)
	q = append(q, start)
	steps[start] = 0

	var head Point2
	for len(q) > 0 {
		head, q = q[0], q[1:]
		if head == target {
			return steps[head]
		}
		ns := steps[head] + 1
		for _, step := range STEPS4 {
			nx, ny := head.x+step[0], head.y+step[1]
			if nx < 0 || ny < 0 || nx >= width || ny >= height {
				continue
			}
			if field[ny][nx] == 0 {
				continue
			}
			np := Point2{nx, ny}
			if ss, ok := steps[np]; !ok || ss > ns {
				steps[np] = ns
				q = append(q, np)
			}
		}
	}

	return -1
}

func traverseAll(field [][]int, markers map[int]Point2, retToBase bool) int {
	mm := make([]int, 0, len(markers))
	var msk uint16
	for marker := range markers {
		mm = append(mm, marker)
		msk |= (1 << marker)
	}
	var recurse func(marker int, visited uint16) int
	recurse = func(marker int, visited uint16) int {
		visited |= (1 << marker)
		if visited == msk {
			if retToBase {
				return findShortestPath(field, markers[marker], markers[0])
			}
			return 0
		}
		minPath := ALOT
		for _, nm := range mm {
			if visited&(1<<nm) > 0 {
				continue
			}
			pp := findShortestPath(field, markers[marker], markers[nm]) + recurse(nm, visited)
			minPath = min(minPath, pp)
		}
		return minPath
	}

	return recurse(0, uint16(0))
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field, markers := makeField(lines)

	printf("markers: %+v", markers)

	println(printIntFieldWithSubs(field, "", map[int]string{
		0: "#",
		1: ".",
	}))

	minPath := traverseAll(field, markers, false)
	printf("min path: %d", minPath)

	minPath2 := traverseAll(field, markers, true)
	printf("min path2: %d", minPath2)
}
