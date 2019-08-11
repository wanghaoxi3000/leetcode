package leetcode

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	var tests = []struct {
		in1 []int
		in2 []int
		out []int
	}{
		{
			[]int{1, 2, 2, 1},
			[]int{2, 2},
			[]int{2},
		},
		{
			[]int{4, 9, 5},
			[]int{9, 4, 9, 8, 4},
			[]int{9, 4},
		},
	}

	for i, item := range tests {
		ret := intersection(item.in1, item.in2)
		if !reflect.DeepEqual(ret, item.out) {
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}
