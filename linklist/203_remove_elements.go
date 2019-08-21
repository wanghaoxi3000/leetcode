/*
203. 移除链表元素

删除链表中等于给定值 val 的所有节点。

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/
package linklist

// 使用虚拟头节点
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{0, head}
	cur := dummyHead

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

// 递归算法
func removeElementsRecursion(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElementsRecursion(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
