package leetcode

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
	var tests = []struct {
		in  [][]int
		out int
	}{
		{[][]int{[]int{0, 0}, []int{1, 0}, []int{2, 0}}, 2},
	}

	for i, item := range tests {
		ret := numberOfBoomerangs(item.in)
		if ret != item.out {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in, ret, item.out)
		}
	}
}
