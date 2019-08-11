package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {
	var tests = []struct {
		in1 []int
		out int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for i, item := range tests {
		ret := missingNumberXor(item.in1)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}

	for i, item := range tests {
		ret := missingNumberSum(item.in1)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}
}
