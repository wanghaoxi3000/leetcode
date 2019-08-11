package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct { // Test table
		in1 []int
		in2 int
		out []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for i, item := range tests {
		ret := twoSum(item.in1, item.in2)
		if !reflect.DeepEqual(ret, item.out) {
			t.Errorf("%d. %v, %v => %d, wanted: %d", i, item.in1, item.in2, ret, item.out)
		}
	}
}
