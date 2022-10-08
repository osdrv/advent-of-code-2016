package main

import (
	"os"
	"sort"
)

// value 53 goes to bot 19
func parseValToBot(s string) (int, int) {
	var val, bid int
	ptr := 0
	_, ptr = readStr(s, ptr, "value ")
	val, ptr = readInt(s, ptr)
	_, ptr = readStr(s, ptr, " goes to bot ")
	bid, ptr = readInt(s, ptr)
	return val, bid
}

// bot 125 gives low to output 4 and high to bot 29
func parseBotToOut(s string) (int, int, int) {
	var bid, out1, out2 int
	ptr := 0
	_, ptr = readStr(s, ptr, "bot ")
	bid, ptr = readInt(s, ptr)
	_, ptr = readStr(s, ptr, " gives low to ")
	if matchStr(s, ptr, "output") {
		_, ptr = readStr(s, ptr, "output ")
		out1, ptr = readInt(s, ptr)
		out1 |= 1 << 31
	} else if matchStr(s, ptr, "bot") {
		_, ptr = readStr(s, ptr, "bot ")
		out1, ptr = readInt(s, ptr)
	}
	_, ptr = readStr(s, ptr, " and high to ")
	if matchStr(s, ptr, "output") {
		_, ptr = readStr(s, ptr, "output ")
		out2, ptr = readInt(s, ptr)
		out2 |= 1 << 31
	} else if matchStr(s, ptr, "bot") {
		_, ptr = readStr(s, ptr, "bot ")
		out2, ptr = readInt(s, ptr)
	}
	return bid, out1, out2
}

const (
	_ int = iota
	VTB
	BTV
)

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	instrs := make([][4]int, 0, len(lines))

	for _, line := range lines {
		if line[:5] == "value" {
			val, bid := parseValToBot(line)
			instrs = append(instrs, [4]int{VTB, val, bid, -1})
			printf("val %d -> bot %d", val, bid)
		} else if line[:3] == "bot" {
			bid, v1, v2 := parseBotToOut(line)
			instrs = append(instrs, [4]int{BTV, bid, v1, v2})
			printf("bot %d -> v_low: %d, v_high: %d", bid, v1, v2)
		} else {
			fatalf("unknown instruction: %s", line)
		}
	}

	printf("instrs: %+v", instrs)

	outs := make(map[int]int)
	bots := make(map[int][]int)

	for i := 0; i < len(instrs); i++ {
		if instrs[i][0] != VTB {
			continue
		}
		val, bid := instrs[i][1], instrs[i][2]
		if _, ok := bots[bid]; !ok {
			bots[bid] = make([]int, 0, 2)
		}
		bots[bid] = append(bots[bid], val)
		sort.Ints(bots[bid])
	}

	printf("bots: %+v", bots)

	cmp1, cmp2 := 17, 61
	//cmp1, cmp2 := 2, 5

	pc := 0
	exc := false
	for {
		if pc >= len(instrs) {
			pc %= len(instrs)
			if !exc {
				break
			}
			exc = false
		}
		if instrs[pc][0] != BTV {
			pc++
			continue
		}
		//printf("instr(%d): %+v", pc, instrs[pc])
		bid, out1, out2 := instrs[pc][1], instrs[pc][2], instrs[pc][3]
		bot := bots[bid]
		if len(bot) < 2 {
			pc++
			continue
		}
		if len(bot) > 2 {
			panic("wtf")
		}

		if bot[0] == cmp1 && bot[1] == cmp2 {
			printf("bot %d is responsible for cmp: %d Vs %d", bid, cmp1, cmp2)
			//break
		}

		exc = true

		if (out1 & (1 << 31)) > 0 {
			// send to out
			dest1 := out1 & (^(1 << 31))
			outs[dest1] = bot[0]
			printf("outs %d gets a value: %d", dest1, bot[0])
		} else {
			bots[out1] = append(bots[out1], bot[0])
			sort.Ints(bots[out1])
		}
		if (out2 & (1 << 31)) > 0 {
			// send to out
			dest2 := out2 & (^(1 << 31))
			outs[dest2] = bot[1]
			printf("outs %d gets a value: %d", dest2, bot[1])
		} else {
			bots[out2] = append(bots[out2], bot[1])
			sort.Ints(bots[out2])
		}
		bots[bid] = make([]int, 0, 2)
		pc++
	}

	printf("outs: %+v", outs)

	printf("mult: %d", outs[0]*outs[1]*outs[2])
}
