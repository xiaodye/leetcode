package leetcode

import "sort"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 收集节点值到切片
	values := []int{}
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	// 排序
	sort.Ints(values)

	// 写回链表
	curr = head
	for i := 0; i < len(values); i++ {
		curr.Val = values[i]
		curr = curr.Next
	}

	return head
}
