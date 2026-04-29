package leetcode

func sortColors(nums []int) {
	i := 0
	j := 0

	// 第一趟：把 0 换到前面
	for j < len(nums) {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}

	// 第二趟：把 1 换到 0 的后面
	j = i
	for j < len(nums) {
		if nums[j] == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
}
