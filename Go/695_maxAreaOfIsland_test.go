package leetcode

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	var tests = []struct { // Test table
		in  [][]int
		out int
	}{
		{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			6,
		},
		{
			[][]int{
				{1},
			},
			1,
		},
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			0,
		},
	}

	for i, item := range tests {
		ret := maxAreaOfIsland(item.in)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in, ret, item.out)
		}
	}
}
