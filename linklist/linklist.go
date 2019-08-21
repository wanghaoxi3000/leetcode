package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateLinkList 根据数组创建链表
func CreateLinkList(vals []int) *ListNode {
	head := &ListNode{0, nil}
	cur := head
	for _, num := range vals {
		cur.Next = &ListNode{num, nil}
		cur = cur.Next
	}
	return head.Next
}

// CompareLinkList 比较两个链表是否相同
func CompareLinkList(SourceHead *ListNode, DestHead *ListNode) bool {
	source := SourceHead
	dest := DestHead
	for source != nil {
		if source.Val != dest.Val {
			return false
		}
		source = source.Next
		dest = dest.Next
	}
	if dest != nil {
		return false
	}
	return true
}

// LinkListConvSlice 链表转换为切片
func LinkListConvSlice(head *ListNode) []int {
	dest := []int{}
	cur := head
	for cur != nil {
		dest = append(dest, cur.Val)
		cur = cur.Next
	}
	return dest
}
