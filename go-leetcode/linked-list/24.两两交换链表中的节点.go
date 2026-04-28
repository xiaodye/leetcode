package leetcode

func swapPairs(head *ListNode) *ListNode {
	// 0 | 1 节点情况
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: head}
	curr := dummy

	for curr.Next != nil && curr.Next.Next != nil {
		node1 := curr.Next
		node2 := curr.Next.Next
		node3 := curr.Next.Next.Next

		// curr -> 第二个节点
		curr.Next = node2
		// 第二个节点 -> 第一个节点
		node2.Next = node1
		// 第一个节点 -> 第三个节点
		node1.Next = node3

		curr = node1
	}

	return dummy.Next
}
