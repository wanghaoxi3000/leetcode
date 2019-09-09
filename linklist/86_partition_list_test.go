package linklist

import (
	"testing"
)

func TestPartition(t *testing.T) {
	var tests = []struct {
		linkListValIn1 []int
		in2            int
		linkListValOut []int
	}{
		{
			[]int{1, 4, 3, 2, 5, 2},
			3,
			[]int{1, 2, 2, 4, 3, 5},
		},
	}

	for i, item := range tests {
		in1 := CreateLinkList(item.linkListValIn1)
		wanted := CreateLinkList(item.linkListValOut)
		ret := partition(in1, item.in2)

		if !CompareLinkList(ret, wanted) {
			t.Errorf("%d. %v %v => %v, wanted: %v", i, item.linkListValIn1, item.in2, LinkListConvSlice(ret), item.linkListValOut)
		}
	}
}
