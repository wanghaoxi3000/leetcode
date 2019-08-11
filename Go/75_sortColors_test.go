package leetcode

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	var tests = []struct {
		in1 []int
		out []int
	}{
		{
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
	}

	for i, item := range tests {
		sortColors(item.in1)
		if !reflect.DeepEqual(item.in1, item.out) {
			t.Errorf("%d. %v, wanted: %v", i, item.in1, item.out)
		}
	}
}
