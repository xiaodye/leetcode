package leetcode

func isPalindrome(head *ListNode) bool {
	// 1. 将链表值存入切片
	values := []int{}
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	// 2. 双指针检查回文
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		if values[i] != values[j] {
			return false
		}
	}
	return true
}
