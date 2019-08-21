/*
24. 两两交换链表中的节点

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
package linklist

func swapPairs(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	cur := dummyHead
	var node1, node2 *ListNode
	for cur.Next != nil && cur.Next.Next != nil {
		node1 = cur.Next
		node2 = node1.Next
		node1.Next = node2.Next
		node2.Next = node1

		cur.Next = node2
		cur = node1
	}
	return dummyHead.Next
}
