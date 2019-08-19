/*
203. 移除链表元素

删除链表中等于给定值 val 的所有节点。

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/
package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

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
