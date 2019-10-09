package tree

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	testTree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	exceptRet := []int{1, 3, 2}

	var testRet []int
	testRet = inorderTraversal(testTree)
	if !reflect.DeepEqual(exceptRet, testRet) {
		t.Errorf("Test inorderTraversal restlt: %v, wanted: %v", testRet, exceptRet)
	}

	testRet = inorderTraversalRecursion(testTree)
	if !reflect.DeepEqual(exceptRet, testRet) {
		t.Errorf("Test preorderTraversal restlt: %v, wanted: %v", testRet, exceptRet)
	}
}
