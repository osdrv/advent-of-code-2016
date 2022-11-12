package main

func popCnt(n uint64) int {
	cnt := 0
	for n != 0 {
		cnt++
		n &= n - 1
	}
	return cnt
}

func isWall(fn, x, y uint64) bool {
	return popCnt(fn+x*x+3*x+2*x*y+y+y*y)%2 == 1
}

func shortestPathTo(fn int, start, dest Point2) int {
	dist := func(p Point2) int {
		return abs(p.x-dest.x) + abs(p.y-dest.y)
	}

	q := NewBinHeap(func(a, b Point2) bool {
		return dist(a) < dist(b)
	})

	steps := make(map[Point2]int)
	enq := make(map[Point2]bool)

	q.Push(start)
	steps[start] = 0
	enq[start] = true

	for q.Size() > 0 {
		head := q.Pop()
		delete(enq, head)
		if head == dest {
			return steps[head]
		}
		for _, step := range STEPS4 {
			nx, ny := head.x+step[0], head.y+step[1]
			if nx < 0 || ny < 0 {
				continue
			}
			if isWall(uint64(fn), uint64(nx), uint64(ny)) {
				continue
			}
			ns := steps[head] + 1
			np := Point2{nx, ny}
			if ps, ok := steps[np]; !ok || ps > ns {
				steps[np] = ns
				if !enq[np] {
					enq[np] = true
					q.Push(np)
				}
			}
		}
	}

	return -1
}

func cntTraverse(fn int, start Point2, rng int) int {
	q := make([]Point2, 0, 1)
	enq := make(map[Point2]bool)

	steps := make(map[Point2]int)

	q = append(q, start)
	steps[start] = 0
	enq[start] = true

	var head Point2
	for len(q) > 0 {
		head, q = q[0], q[1:]
		delete(enq, head)
		if steps[head] >= 50 {
			continue
		}
		for _, step := range STEPS4 {
			nx, ny := head.x+step[0], head.y+step[1]
			if nx < 0 || ny < 0 {
				continue
			}
			if isWall(uint64(fn), uint64(nx), uint64(ny)) {
				continue
			}
			np := Point2{nx, ny}
			ns := steps[head] + 1
			if ps, ok := steps[np]; !ok || ps > ns {
				steps[np] = ns
				if !enq[np] {
					enq[np] = true
					q = append(q, np)
				}
			}
		}
	}

	return len(steps)
}

func main() {
	//FN := 10
	//DEST := Point2{7, 4}
	FN := 1350
	START := Point2{1, 1}
	DEST := Point2{31, 39}

	dist := shortestPathTo(FN, START, DEST)
	printf("shortest path to %+v: %d", DEST, dist)

	rng := cntTraverse(FN, START, 50)
	printf("total reachable area: %d", rng)
}
