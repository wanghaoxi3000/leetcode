package tree

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	testTree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}
	exceptRet := []int{1, 2, 3}

	var testRet []int
	testRet = preorderTraversalRecursion(testTree)
	if !reflect.DeepEqual(exceptRet, testRet) {
		t.Errorf("Test preorderTraversalRecursion restlt: %v, wanted: %v", testRet, exceptRet)
	}

	testRet = preorderTraversal(testTree)
	if !reflect.DeepEqual(exceptRet, testRet) {
		t.Errorf("Test preorderTraversal restlt: %v, wanted: %v", testRet, exceptRet)
	}
}
