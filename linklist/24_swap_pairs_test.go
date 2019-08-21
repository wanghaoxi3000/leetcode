package linklist

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	var tests = []struct {
		linkListValIn  []int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 4, 3},
		},
	}

	for i, item := range tests {
		source := CreateLinkList(item.linkListValIn)
		ret := swapPairs(source)
		dest := CreateLinkList(item.linkListValOut)
		if !CompareLinkList(ret, dest) {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.linkListValIn, LinkListConvSlice(ret), item.linkListValOut)
		}
	}
}
