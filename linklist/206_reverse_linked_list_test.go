package linklist

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	var tests = []struct {
		linkListValIn  []int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
	}

	for i, item := range tests {
		source := CreateLinkList(item.linkListValIn)
		dest := CreateLinkList(item.linkListValOut)
		reverseHead := reverseList(source)
		if !CompareLinkList(reverseHead, dest) {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.linkListValIn, LinkListConvSlice(reverseHead), item.linkListValOut)
		}
	}
}

func TestReverseListRecursion(t *testing.T) {
	var tests = []struct {
		linkListValIn  []int
		linkListValOut []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
	}

	for i, item := range tests {
		source := CreateLinkList(item.linkListValIn)
		dest := CreateLinkList(item.linkListValOut)
		reverseHead := reverseListRecursion(source)
		if !CompareLinkList(reverseHead, dest) {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.linkListValIn, LinkListConvSlice(reverseHead), item.linkListValOut)
		}
	}
}
