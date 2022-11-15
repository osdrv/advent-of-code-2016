package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseRange(s string) (uint64, uint64) {
	ss := strings.SplitN(s, "-", 2)
	low, err := strconv.ParseUint(ss[0], 10, 64)
	noerr(err)
	high, err := strconv.ParseUint(ss[1], 10, 64)
	noerr(err)
	return low, high
}

func compact(ranges [][2]uint64) [][2]uint64 {
	res := make([][2]uint64, 0, 1)
	ptr := 0
	for ptr < len(ranges) {
		from := ptr
		high := ranges[ptr][1]
		for ptr < len(ranges) && high >= (max(ranges[ptr][0], 1)-1) {
			high = max(ranges[ptr][1], high)
			ptr++
		}
		res = append(res, [2]uint64{ranges[from][0], high})
	}
	return res
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	ranges := make([][2]uint64, 0, len(lines))
	for _, line := range lines {
		low, high := parseRange(line)
		assert(low < high, "low < high")
		ranges = append(ranges, [2]uint64{low, high})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0] || ((ranges[i][0] == ranges[j][0]) && (ranges[i][1] < ranges[j][1]))
	})

	printf("ranges: %+v", ranges)

	ranges = compact(ranges)
	ranges = append(ranges, [2]uint64{4294967295 + 1, 4294967295 + 1})

	printf("compacted: %+v", ranges)

	printf("the min int: %d", ranges[0][1]+1)

	allowed := uint64(0)
	for ptr := 1; ptr < len(ranges); ptr++ {
		allowed += (ranges[ptr][0] - ranges[ptr-1][1] - 1)
	}
	printf("allowed IPs: %d", allowed)
}
