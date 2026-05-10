package leetcode

/*
* Do not return anything, modify nums in-place instead.
 */
func rotate(nums []int, k int) {
	k = k % len(nums)

	if k == 0 {
		return
	}

	// 反转整个数组
	reverse(nums, 0, len(nums)-1)
	// 反转前 k 个
	reverse(nums, 0, k-1)
	// 反转剩余
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
