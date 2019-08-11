package leetcode

import (
	"testing"
)

func TestFindLengthOfLCIS(t *testing.T) {
	var tests = []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 3, 5, 4, 7},
			3,
		},
		{
			[]int{2, 2, 2, 2, 2},
			1,
		},
		{
			[]int{1, 3, 5, 7},
			4,
		},
	}

	for i, item := range tests {
		ret := findLengthOfLCIS(item.in)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in, ret, item.out)
		}
	}
}
