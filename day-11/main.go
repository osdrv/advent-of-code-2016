package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ObjectType uint8

const (
	_ ObjectType = iota
	Generator
	Microchip
)

type Object struct {
	name string
	typ  ObjectType
}

func (o Object) Handle() string {
	s := strings.ToUpper(o.name)[0:1]
	if o.typ == Generator {
		s += "G"
	} else {
		s += "M"
	}
	return s
}

var (
	rr = regexp.MustCompile("([\\w\\-]+\\smicrochip)|([\\w]+\\sgenerator)")
)

func parseLine(s string) []Object {
	res := make([]Object, 0, 1)
	for _, mtch := range rr.FindAllStringSubmatch(s, -1) {
		mm := strings.SplitN(mtch[0], " ", 2)
		var typ ObjectType
		switch mm[1] {
		case "generator":
			typ = Generator
		case "microchip":
			typ = Microchip
		default:
			panic("wtf")
		}
		if typ == Microchip {
			mmm := strings.SplitN(mm[0], "-", 2)
			mm[0] = mmm[0]
		}
		res = append(res, Object{mm[0], typ})
	}
	return res
}

func printFloorMap(floors []uint16, floor int, objs map[int]Object) string {
	var b bytes.Buffer

	for i := len(floors) - 1; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("F%d ", i+1))
		if i == floor {
			b.WriteByte('E')
		} else {
			b.WriteByte('.')
		}
		b.WriteString("  ")
		for j := 0; j < len(objs)/2; j++ {
			if floors[i]&uint16(1<<(2*j)) > 0 {
				b.WriteString(objs[2*j].Handle())
			} else {
				b.WriteString(". ")
			}
			b.WriteByte(' ')
			if floors[i]&uint16(1<<(2*j+1)) > 0 {
				b.WriteString(objs[2*j+1].Handle())
			} else {
				b.WriteString(". ")
			}
			b.WriteByte(' ')
		}
		b.WriteByte('\n')
	}

	return b.String()
}

func parseData(lines []string) ([]uint16, map[int]Object) {
	levels := make([]uint16, 0, len(lines))
	binding := make(map[string]int)
	objs := make(map[int]Object)
	cnt := 0

	for _, line := range lines {
		var level uint16
		for _, obj := range parseLine(line) {
			if _, ok := binding[obj.name]; !ok {
				binding[obj.name] = cnt
				cnt++
			}
			p := binding[obj.name] * 2
			if obj.typ == Microchip {
				p += 1
			}
			objs[p] = obj
			level |= (1 << p)
		}
		levels = append(levels, level)
	}

	return levels, objs
}

func copyFloors(floors []uint16) []uint16 {
	cp := make([]uint16, len(floors))
	copy(cp, floors)
	return cp
}

const (
	GENS  uint16 = 0b0101010101010101
	CHIPS uint16 = 0b1010101010101010
)

func isComplete(floors []uint16) bool {
	for i := 0; i < len(floors)-1; i++ {
		if floors[i] != 0 {
			return false
		}
	}
	return (floors[len(floors)-1] & GENS) == ((floors[len(floors)-1] & CHIPS) >> 1)
}

func isValid(floor uint16) bool {
	/*
		no objects at all = valid
		generators only = valid
		chips only = valid
		paired generators-chip sets + chips = valid
		any outstanding generator and an unpaired chip = invalid
	*/
	if floor&CHIPS == 0 || // all-gens
		floor&GENS == 0 { // all-chips
		return true
	}

	return ((((floor & CHIPS) >> 1) ^ (floor & GENS)) & ((floor & CHIPS) >> 1)) == 0

	// a valid floor: HG HM LG
	// 1  1  1  0 -> 0b0111
	// CHIPS      -> 0b0010
	// GENS       -> 0b0101
	// (CHIPS>>1)&GENS -> 0b0001 & 0b0101 -> 0b0001

	//(((floor & CHIPS) >> 1) ^ (floor & GENS))&(floor & CHIPS) >> 1) == 0,
}

func getAllMasks(floor uint16) []uint16 {
	res := make([]uint16, 0, 1)
	for i := 0; i < 16; i++ {
		if floor&(1<<i) == 0 {
			continue
		}
		res = append(res, (1 << i))
		for j := i + 1; j < 16; j++ {
			if floor&(1<<j) == 0 {
				continue
			}
			res = append(res, (1<<i)|(1<<j))
		}
	}
	return res
}

func liftUp(floors []uint16, objs map[int]Object) int {
	type queueItem struct {
		floor int
		state []uint16
	}

	expp := floors[0] ^ floors[1] ^ floors[2] ^ floors[3]

	steps := make(map[uint64]int)

	debugf("expect: %016b", expp)

	q := make([]queueItem, 0, 1)
	q = append(q, queueItem{0, copyFloors(floors)})
	steps[mkKey(floors, 0)] = 0

	var head queueItem
	for len(q) > 0 {
		head, q = q[0], q[1:]
		key := mkKey(head.state, head.floor)
		if DEBUG {
			debugf("steps: %d", steps[key])
			debugf("key: 0b%064b", key)
			println(printFloorMap(head.state, head.floor, objs))
		}
		if isComplete(head.state) {
			return steps[key]
		}
		for _, df := range []int{1, -1} {
			// we assume the total amount of pairs is less or eql to 8
			nFloor := head.floor + df
			if nFloor < 0 || nFloor >= len(head.state) {
				continue
			}
			debugf("all masks for %016b: %+v", head.state[head.floor], getAllMasks(head.state[head.floor]))
			for _, msk := range getAllMasks(head.state[head.floor]) {
				debugf("mask: %016b", msk)
				ncs := head.state[head.floor] & (^msk)
				nns := head.state[nFloor] | (msk & head.state[head.floor])
				debugf("old cur state: %016b", head.state[head.floor])
				debugf("new cur state: %016b", ncs)
				debugf("old new state: %016b", head.state[nFloor])
				debugf("new new state: %016b", nns)
				if isValid(ncs) && isValid(nns) {
					statecp := copyFloors(head.state)
					statecp[head.floor] = ncs
					statecp[nFloor] = nns
					ns := steps[mkKey(head.state, head.floor)] + 1
					k := mkKey(statecp, nFloor)
					if s, ok := steps[k]; !ok || s > ns {
						steps[k] = ns
						q = append(q, queueItem{floor: nFloor, state: statecp})
					}
				}
			}
		}
	}

	return -1
}

func mkKey(floors []uint16, floor int) uint64 {
	return uint64(floors[0]) | (uint64(floors[1]) << 16) | (uint64(floors[2]) << 32) | (uint64(floors[3]) << 48) | ((uint64(floor) & 0b11) << 62)
}

func main() {
	f, err := os.Open("INPUT.2")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	printf("file data: %+v", lines)

	floors, objs := parseData(lines)
	debugf("floors: %+v", floors)
	debugf("objs: %+v", objs)

	for ix := range floors {
		debugf("%s", lines[ix])
		debugf("%016b", floors[ix])
	}

	minMoves := liftUp(floors, objs)
	printf("min moves: %d", minMoves)
}
