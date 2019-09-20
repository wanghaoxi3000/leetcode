package linklist

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		linkListValIn1 []int
		in2            int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 2, 3, 5},
		},
	}

	for i, item := range tests {
		in1 := CreateLinkList(item.linkListValIn1)
		wanted := CreateLinkList(item.linkListValOut)
		ret := removeNthFromEnd(in1, item.in2)

		if !CompareLinkList(ret, wanted) {
			t.Errorf("%d. %v %v => %v, wanted: %v", i, item.linkListValIn1, item.in2, LinkListConvSlice(ret), item.linkListValOut)
		}
	}
}
