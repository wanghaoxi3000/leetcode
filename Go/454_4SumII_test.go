package leetcode

import (
	"testing"
)

func TestFourSumCount(t *testing.T) {
	var tests = []struct {
		in1 []int
		in2 []int
		in3 []int
		in4 []int
		out int
	}{
		{
			[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2,
		},
		{
			[]int{-1, 1, 1, 1, -1}, []int{0, -1, -1, 0, 1}, []int{-1, -1, 1, -1, -1}, []int{0, 1, 0, -1, -1}, 132,
		},
	}

	for i, item := range tests {
		ret := fourSumCount(item.in1, item.in2, item.in3, item.in4)
		if ret != item.out {
			t.Errorf("%d. %v, %v, %v, %v => %v, wanted: %v",
				i, item.in1, item.in2, item.in3, item.in4, ret, item.out)
		}
	}
}
