package tree

/*
94. Binary Tree Inorder Traversal 二叉树的中序遍历
给定一个二叉树，返回它的中序遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/* inorderTraversal 迭代中序遍历 左中右
1. 辅助结点p初始化为根结点，while循环的条件是栈不为空或者辅助结点p不为空
2. 在循环中首先判断如果辅助结点p存在，那么先将p加入栈中，此时p指向其左子结点
3. 否则如果p不存在的话，表明没有左子结点，我们取出栈顶结点，将p的结点值加入结果ret中，将p指向栈顶结点的右子结点
*/
func inorderTraversal(root *TreeNode) []int {
	p := root
	ret := []int{}

	stack := []*TreeNode{} // 存储值模拟栈的列表
	stackTop := 0          // 栈顶标记

	for stackTop > 0 || p != nil {
		if p != nil {
			stack = append(stack, p)
			stackTop++
			p = p.Left
		} else {
			stackTop--
			t := stack[stackTop]
			stack = stack[0:stackTop]
			ret = append(ret, t.Val)
			p = t.Right
		}
	}

	return ret
}

// inorderTraversalRecursion 递归中序遍历
func inorderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	ret = append(ret, inorderTraversalRecursion(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversalRecursion(root.Right)...)

	return ret
}
