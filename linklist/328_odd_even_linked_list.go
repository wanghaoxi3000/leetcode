/*
328. 奇偶链表 Odd Even Linked List
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

示例 2:
输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL

说明:
应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/
package linklist

/* oddEvenList
1. 设定两个虚拟节点，dummyHead1用来保存奇节点，dummyHead2来保存偶节点；
2. 遍历整个原始链表，将奇节点放于dummyHead1中，其余的放置在dummyHead2中
3. 遍历结束后，将dummyHead2插入到dummyHead1后面
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	dummyHead1 := new(ListNode)
	dummyHead2 := new(ListNode)
	p1 := dummyHead1
	p2 := dummyHead2
	p := head

	for i := 0; p != nil; i++ {
		if i%2 == 0 {
			p1.Next = p
			p1 = p1.Next
			p = p.Next
			p1.Next = nil
		} else {
			p2.Next = p
			p2 = p2.Next
			p = p.Next
			p2.Next = nil
		}
	}

	p1.Next = dummyHead2.Next
	return head
}
