package linklist

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	var tests = []struct {
		sourceListVal  []int
		removeNodeVal  int
		removedListVal []int
	}{
		{
			[]int{4, 5, 1, 9},
			1,
			[]int{4, 5, 9},
		},
		{
			[]int{4, 5, 1, 9},
			5,
			[]int{4, 1, 9},
		},
	}

	for i, item := range tests {
		source := CreateLinkList(item.sourceListVal)
		dest := CreateLinkList(item.removedListVal)
		removeNode := source
		for removeNode.Val != item.removeNodeVal {
			removeNode = removeNode.Next
		}
		deleteNode(removeNode)
		if !CompareLinkList(source, dest) {
			t.Errorf("%d. %v => %v, wanted: %v", i, item.sourceListVal, LinkListConvSlice(source), item.removedListVal)
		}
	}
}
