package leetcode

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var tests = []struct {
		in1 int
		in2 []int
		out int
	}{
		{
			7,
			[]int{2, 3, 1, 2, 4, 3},
			2,
		},
	}

	for i, item := range tests {
		ret := minSubArrayLen(item.in1, item.in2)
		if ret != item.out {
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}
