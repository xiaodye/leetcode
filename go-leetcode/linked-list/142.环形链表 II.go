package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// 相遇，找环入口
			x := head
			y := slow

			for x != y {
				x = x.Next
				y = y.Next
			}

			return x
		}
	}

	return nil
}
