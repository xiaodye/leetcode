package leetcode

func findMin(nums []int) int {
	// 旋转后，还是原数组情况，元素升序
	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		// 前半段无序，忽略后半段
		if nums[0] <= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
