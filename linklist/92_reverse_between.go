/*
Reverse Linked List II 反转链表 II

反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/
package linklist

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	cur := dummy
	var front *ListNode
	for i := 1; i < m; i++ {
		cur = cur.Next
	}
	pre := cur
	last := cur.Next

	for i := m; i <= n; i++ {
		cur = pre.Next
		pre.Next = cur.Next
		cur.Next = front
		front = cur
	}

	cur = pre.Next
	pre.Next = front
	last.Next = cur
	return dummy.Next
}
