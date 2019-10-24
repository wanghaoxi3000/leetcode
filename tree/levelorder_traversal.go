package tree

/*
102. Binary Tree Level Order Traversal 二叉树的层次遍历

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]map[*TreeNode]int, 0)
	queue = append(queue, map[*TreeNode]int{root: 0})

	for len(queue) != 0 {
		for k, v := range queue[0] {
			queue = queue[1:]
			if v == len(res) {
				res = append(res, []int{})
			}

			res[v] = append(res[v], k.Val)
			node := k
			if node.Left != nil {
				queue = append(queue, map[*TreeNode]int{node.Left: v + 1})
			}
			if node.Right != nil {
				queue = append(queue, map[*TreeNode]int{node.Right: v + 1})
			}
		}
	}

	return res
}
