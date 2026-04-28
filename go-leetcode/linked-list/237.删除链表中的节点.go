package leetcode

func deleteNode(node *ListNode) {
	// 我们只需要删除 next 这个节点，并且让 node.val = node.next
	//
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
