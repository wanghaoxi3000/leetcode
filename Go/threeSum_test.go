package leetcode

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var tests = []struct { // Test table
		in1 []int
		out [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for i, item := range tests {
		ret := threeSum(item.in1)
		if !reflect.DeepEqual(ret, item.out) {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.in1, ret, item.out)
		}
	}
}
