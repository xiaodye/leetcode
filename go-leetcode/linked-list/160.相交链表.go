package leetcode

func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)

	temp := headA
	for temp != nil {
		visited[temp] = true
		temp = temp.Next
	}

	temp = headB
	for temp != nil {
		if visited[temp] {
			return temp
		}
		temp = temp.Next
	}

	return nil
}
