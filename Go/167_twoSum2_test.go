package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	var tests = []struct {
		in1 []int
		in2 int
		out []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	}

	for i, item := range tests {
		ret := twoSum2(item.in1, item.in2)
		if !reflect.DeepEqual(ret, item.out) {
			t.Errorf("%d. %v %v => %v, wanted: %v", i, item.in1, item.in2, ret, item.out)
		}
	}
}
