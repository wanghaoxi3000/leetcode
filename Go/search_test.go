package leetcode

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct { // Test table
		in1 []int
		in2 int
		out int
	}{
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			[]int{1, 3},
			3,
			1,
		},
		{
			[]int{8, 1, 2, 3, 4, 5, 6, 7},
			6,
			6,
		},
	}

	for i, item := range tests {
		ret := search(item.in1, item.in2)
		if ret != item.out {
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}
