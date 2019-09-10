package linklist

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
	var tests = []struct {
		linkListValIn1 []int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 3, 5, 2, 4},
		},
		{
			[]int{2, 1, 3, 5, 6, 4, 7},
			[]int{2, 3, 6, 7, 1, 5, 4},
		},
	}

	for i, item := range tests {
		in1 := CreateLinkList(item.linkListValIn1)
		wanted := CreateLinkList(item.linkListValOut)
		ret := oddEvenList(in1)

		if !CompareLinkList(ret, wanted) {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.linkListValIn1, LinkListConvSlice(ret), item.linkListValOut)
		}
	}
}
