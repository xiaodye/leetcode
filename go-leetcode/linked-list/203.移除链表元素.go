package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	// 定义一个虚拟头结点
	dummy := &ListNode{Next: head}
	curr := dummy

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}
