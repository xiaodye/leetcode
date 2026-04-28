package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		n1 := 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		n2 := 0
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry == 1 {
		curr.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}
