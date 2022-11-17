package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	x, y int

	size, used, avail, use int
}

var re = regexp.MustCompile("\\/dev\\/grid\\/node-x(\\d+)-y(\\d+)\\s+(\\d+)T\\s+(\\d+)T\\s+(\\d+)T\\s+(\\d+)%")

func parseNode(s string) *Node {
	mtch := re.FindStringSubmatch(s)
	return &Node{
		x:     parseInt(mtch[1]),
		y:     parseInt(mtch[2]),
		size:  parseInt(mtch[3]),
		used:  parseInt(mtch[4]),
		avail: parseInt(mtch[5]),
		use:   parseInt(mtch[6]),
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("Node{x: %d, y: %d, size: %d, used: %d, avail: %d, use: %d}", n.x, n.y, n.size, n.used, n.avail, n.use)
}

func isViablePair(n1, n2 *Node) bool {
	return n1.used > 0 && n1.used <= n2.avail
}

func printNodes(nodes map[Point2]*Node, maxx, maxy int) string {
	var b bytes.Buffer
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			node := nodes[Point2{x, y}]
			if node.used == 0 {
				b.WriteByte('0')
			} else if node.used > 100 {
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
			b.WriteByte(' ')
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func computeShortestPath(nodes map[Point2]*Node, start, target Point2, maxx, maxy int) int {
	q := make([]Point2, 0, 1)
	q = append(q, start)

	steps := make(map[Point2]int)
	steps[start] = 0

	var head Point2
	for len(q) > 0 {
		head, q = q[0], q[1:]
		if head == target {
			return steps[head]
		}
		for _, step := range STEPS4 {
			nx, ny := head.x+step[0], head.y+step[1]
			if nx < 0 || ny < 0 || nx > maxx || ny > maxy {
				continue
			}
			np := Point2{nx, ny}
			if nodes[head].used > nodes[np].size {
				continue
			}
			ns := steps[head] + 1
			if s, ok := steps[np]; !ok || ns < s {
				steps[np] = ns
				q = append(q, np)
			}
		}
	}

	return -1
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	lines = lines[2:]

	maxx, maxy := -ALOT, -ALOT
	nodes := make(map[Point2]*Node, len(lines))
	zerox, zeroy := 0, 0
	for _, line := range lines {
		node := parseNode(line)
		if node.used == 0 {
			// I checked the data and I know there would be only 1 empty node
			zerox, zeroy = node.x, node.y
			debugf("found zero node: %+v", *node)
		}
		maxx, maxy = max(maxx, node.x), max(maxy, node.y)
		nodes[Point2{node.x, node.y}] = node
	}

	cntPairs := 0
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1 == n2 {
				continue
			}
			if isViablePair(n1, n2) {
				cntPairs++
			}
		}
	}

	printf("number of viable pairs: %d", cntPairs)

	println(printNodes(nodes, maxx, maxy))

	// 0. find the zero node (there is only 1)
	debugf("zero node position: %d, %d", zerox, zeroy)

	// 1. ensure max used is below the min size

	maxused, minsize := -ALOT, ALOT
	for x := 0; x <= maxx; x++ {
		n0 := nodes[Point2{x, 0}]
		n1 := nodes[Point2{x, 0}]
		maxused = max(max(maxused, n0.used), n1.used)
		minsize = min(min(minsize, n0.size), n1.size)
	}
	printf("max used: %d, min size: %d", maxused, minsize)
	assert(maxused <= minsize, "max used is under the min size")

	// 2. navigate empty node to the upper right corner

	/*

			X . . . . . . . _ G
			. . . . . . . . . . 0

			X . . . . . . . G _
			. . . . . . . . . . 1

			X . . . . . . . G .
			. . . . . . . . . _ 2

			X . . . . . . . G .
			. . . . . . . . _ . 3

			X . . . . . . . G .
			. . . . . . . _ . . 4

			X . . . . . . _ G .
			. . . . . . . . . . 5

			X . . . . . . G _ .
			. . . . . . . . . . 6

			X . . . . . G _ . .
			. . . . . . . . . . 11

			X . . . . G _ . . .
			. . . . . . . . . . 16

			X . . . G _ . . . .
			. . . . . . . . . . 21

			X . . G _ . . . . .
			. . . . . . . . . . 26

			X . G _ . . . . . .
			. . . . . . . . . . 31

			X G _ . . . . . . .
			. . . . . . . . . . 36

		    G _ . . . . . . . .
			. . . . . . . . . . 41

	*/

	steps := computeShortestPath(nodes, Point2{zerox, zeroy}, Point2{maxx - 1, 0}, maxx, maxy)
	printf("min steps: %d", steps)

	steps += 1

	steps += (maxx - 1) * 5

	printf("total steps: %d", steps)

	/*

			initial position:

			(.) _  G
		     .  .  .
		     #  .  .

			move 1:
			(.) G  _
		     .  .  .
		     #  .  .

			 ...

			 (G) _  .
		      .  .  .
		      #  .  .

			  +5 steps in total for a single position move

	*/

	//totSteps := steps +

	// 3. count amnount of 5-cycles

}
