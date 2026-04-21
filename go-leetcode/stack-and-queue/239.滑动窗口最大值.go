package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	deque := []int{} // 存储索引

	for i := 0; i < len(nums); i++ {
		// 保持队列单调递减：移除队尾所有小于当前值的索引
		for len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		// 移除窗口外的索引
		if deque[0] <= i-k {
			deque = deque[1:] // shift
		}

		// 当窗口形成时，记录最大值
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
