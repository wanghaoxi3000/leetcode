/*
86. 分隔链表

给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/
package linklist

/*
1. 设定两个虚拟节点，dummyHead1用来保存小于于该值的链表，dummyHead2来保存大于等于该值的链表
2. 遍历整个原始链表，将小于该值的放于dummyHead1中，其余的放置在dummyHead2中
3. 遍历结束后，将dummyHead2插入到dummyHead1后面
*/
func partition(head *ListNode, x int) *ListNode {
	dummyHead1 := new(ListNode)
	dummyHead2 := new(ListNode)
	prev1 := dummyHead1
	prev2 := dummyHead2

	for cur := head; cur != nil; {
		if cur.Val < x {
			prev1.Next = cur
			prev1 = prev1.Next
			cur = cur.Next
			prev1.Next = nil
		} else {
			prev2.Next = cur
			prev2 = prev2.Next
			cur = cur.Next
			prev2.Next = nil
		}
	}

	prev1.Next = dummyHead2.Next
	return dummyHead1.Next
}
