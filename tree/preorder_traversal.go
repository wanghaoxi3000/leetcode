package tree

/*
144. Binary Tree Preorder Traversal 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/* preorderTraversal 迭代前序遍历
1. 辅助结点p初始化为根结点，while循环的条件是栈不为空或者辅助结点p不为空
2. 在循环中首先判断如果辅助结点p存在，那么先将p加入栈中，然后将p的结点值加入结果res中，此时p指向其左子结点
3. 否则如果p不存在的话，表明没有左子结点，我们取出栈顶结点，将p指向栈顶结点的右子结点
*/
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	stack := []*TreeNode{}
	stackTop := 0
	p := root

	for stackTop > 0 || p != nil {
		if p != nil {
			stack = append(stack, p)
			stackTop++
			ret = append(ret, p.Val)
			p = p.Left
		} else {
			stackTop--
			t := stack[stackTop]
			stack = stack[0:stackTop]
			p = t.Right
		}
	}

	return ret
}

// preorderTraversalRecursion 递归前序遍历
func preorderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{root.Val}
	ret = append(ret, preorderTraversalRecursion(root.Left)...)
	ret = append(ret, preorderTraversalRecursion(root.Right)...)

	return ret
}
