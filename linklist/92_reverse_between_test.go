package linklist

import (
	"testing"
)

func TestReverseBetween(t *testing.T) {
	var tests = []struct {
		linkListValIn1 []int
		in2            int
		in3            int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			4,
			[]int{1, 4, 3, 2, 5},
		},
	}

	for i, item := range tests {
		in1 := CreateLinkList(item.linkListValIn1)
		wanted := CreateLinkList(item.linkListValOut)
		ret := reverseBetween(in1, item.in2, item.in3)

		if !CompareLinkList(ret, wanted) {
			t.Errorf("%d. %v %v %v => %v, wanted: %v", i, item.linkListValIn1, item.in2, item.in3, LinkListConvSlice(ret), item.linkListValOut)
		}
	}
}
