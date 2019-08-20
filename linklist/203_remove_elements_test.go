package linklist

import (
	"testing"
)

func CreateLinkList(vals []int) *ListNode {
	head := &ListNode{0, nil}
	cur := head
	for _, num := range vals {
		cur.Next = &ListNode{num, nil}
		cur = cur.Next
	}
	return head.Next
}

func CompareLinkList(SourceHead *ListNode, DestHead *ListNode) bool {
	source := SourceHead
	dest := DestHead
	for source != nil {
		if source.Val != dest.Val {
			return false
		}
		source = source.Next
		dest = dest.Next
	}
	if dest != nil {
		return false
	}
	return true
}

func LinkListConvList(head *ListNode) []int {
	dest := []int{}
	cur := head
	for cur != nil {
		dest = append(dest, cur.Val)
		cur = cur.Next
	}
	return dest
}

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
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.linkListValIn, item.removeVal, LinkListConvList(ret), item.linkListValOut)
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
			t.Errorf("%d. %v, %v => %v, wanted: %v", i, item.linkListValIn, item.removeVal, LinkListConvList(ret), item.linkListValOut)
		}
	}
}
