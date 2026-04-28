package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 边界情况：两个链表都为空
	if list1 == nil && list2 == nil {
		return nil
	}

	dummy := &ListNode{Val: -1}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	// 连接剩余部分
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}
