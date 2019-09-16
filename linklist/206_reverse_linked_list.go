/*
206. 反转链表 Reverse Linked List
反转一个单链表。

示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

package linklist

/*
1. 每次查看cur节点是否为NULL，如果是，则结束循环，获得结果
2. 如果cur节点不是为NULL，则先设置临时变量next为cur的下一个节点
3. 让cur的下一个节点变成指向pre，而后pre移动cur，cur移动到next
4. 重复1 2 3

时间复杂度 O(n)
空间复杂度 O(1)
*/
func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	cur := head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

/*
递归的方式反转链表
时间复杂度 O(n)
空间复杂度 O(1)
*/
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rhead := reverseListRecursion(head.Next)

	head.Next.Next = head
	head.Next = nil

	return rhead
}
