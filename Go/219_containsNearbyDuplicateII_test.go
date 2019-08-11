package leetcode

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	var tests = []struct {
		in1 []int
		in2 int
		out bool
	}{
		{
			[]int{1, 2, 3, 1},
			3,
			true,
		},
		{
			[]int{1, 0, 1, 1},
			1,
			true,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			2,
			false,
		},
	}

	for i, item := range tests {
		ret := containsNearbyDuplicate(item.in1, item.in2)
		if ret != item.out {
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}
