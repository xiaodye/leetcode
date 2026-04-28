package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	slow, fast := dummy, dummy

	// fast 先走 n 步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// fast 再提前走一步，让 slow 指向待删节点的前驱
	fast = fast.Next

	// 同时移动
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除 slow.Next
	slow.Next = slow.Next.Next

	return dummy.Next
}
