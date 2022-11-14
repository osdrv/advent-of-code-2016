package main

import (
	"bytes"
	"strings"
)

func regenTile(init string, rows int) []string {
	tile := make([]string, 0, rows)
	tile = append(tile, init)
	for i := 1; i < rows; i++ {
		var b bytes.Buffer
		for j := 0; j < len(init); j++ {
			var l, r byte = '.', '.'
			var t byte = '.'
			if j > 0 {
				l = tile[i-1][j-1]
			}
			if j < len(init)-1 {
				r = tile[i-1][j+1]
			}
			if (l == '.' && r == '^') || (l == '^' && r == '.') {
				t = '^'
			}
			b.WriteByte(t)
		}
		tile = append(tile, b.String())
	}
	return tile
}

func cntSafeTiles(tile []string) int {
	cnt := 0
	for i := 0; i < len(tile); i++ {
		for j := 0; j < len(tile[0]); j++ {
			if tile[i][j] == '.' {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	inputs := []struct {
		row    string
		height int
	}{
		{"..^^.", 3},
		{".^^.^.^^^^", 10},
		{".^^^.^.^^^^^..^^^..^..^..^^..^.^.^.^^.^^....^.^...^.^^.^^.^^..^^..^.^..^^^.^^...^...^^....^^.^^^^^^^", 40},
		{".^^^.^.^^^^^..^^^..^..^..^^..^.^.^.^^.^^....^.^...^.^^.^^.^^..^^..^.^..^^^.^^...^...^^....^^.^^^^^^^", 400000},
	}

	for _, ii := range inputs {
		tile := regenTile(ii.row, ii.height)
		debugf(strings.Join(tile, "\n"))
		printf("Safe tiles: %d", cntSafeTiles(tile))
	}
}
