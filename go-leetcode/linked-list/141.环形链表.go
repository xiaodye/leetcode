package leetcode

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)

	curr := head

	for curr != nil {
		if visited[curr] {
			return true
		}
		visited[curr] = true
		curr = curr.Next
	}

	return false
}
