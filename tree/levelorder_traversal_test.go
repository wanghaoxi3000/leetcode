package tree

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	testTree := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	exceptRet := [][]int{
		[]int{3},
		[]int{9, 20},
		[]int{15, 7},
	}

	var testRet [][]int
	testRet = levelOrder(testTree)
	if !reflect.DeepEqual(exceptRet, testRet) {
		t.Errorf("Test levelOrder restlt: %v, wanted: %v", testRet, exceptRet)
	}
}
