package linklist

/*
19. Remove Nth Node From End of List 删除链表的倒数第N个节点

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.

说明：
给定的 n 保证是有效的。

进阶：
你能尝试使用一趟扫描实现吗？
*/

/*
可以设想假设设定了双指针p和q的话，当q指向末尾的NULL，p与q之间相隔的元素个数为n时，那么删除掉p的下一个指针就完成了要求。
1. 设置虚拟节点dummyHead指向head
2. 设定双指针p和q，初始都指向虚拟节点dummyHead
3. 移动q，直到p与q之间相隔的元素个数为n
4. 同时移动p与q，直到q指向的为NULL
5. 将p的下一个节点指向下下个节点
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	p := dummyHead
	q := dummyHead

	for i := 0; i < n+1; i++ {
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next
	return dummyHead.Next
}
