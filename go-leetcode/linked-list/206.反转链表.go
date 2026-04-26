package leetcode

func reverseList(head *ListNode) *ListNode {
	// 0 | 1 个节点
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode = nil
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}

	return pre
}
