package linklist

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	var tests = []struct {
		linkListValIn  []int
		removeVal      int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 6, 3, 4, 5, 6},
			6,
			[]int{1, 2, 3, 4, 5},
		},
	}

	for i, item := range tests {
		source := CreateLinkList(item.linkListValIn)
		ret := removeElements(source, item.removeVal)
		dest := CreateLinkList(item.linkListValOut)
		if !CompareLinkList(ret, dest) {
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.linkListValIn, item.removeVal, LinkListConvSlice(ret), item.linkListValOut)
		}
	}
}

func TestRemoveElementsRecursion(t *testing.T) {
	var tests = []struct {
		linkListValIn  []int
		removeVal      int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 6, 3, 4, 5, 6},
			6,
			[]int{1, 2, 3, 4, 5},
		},
	}

	for i, item := range tests {
		source := CreateLinkList(item.linkListValIn)
		ret := removeElementsRecursion(source, item.removeVal)
		dest := CreateLinkList(item.linkListValOut)
		if !CompareLinkList(ret, dest) {
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.linkListValIn, item.removeVal, LinkListConvSlice(ret), item.linkListValOut)
		}
	}
}
