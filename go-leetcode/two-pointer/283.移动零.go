package leetcode

// Do not return anything, modify nums in-place instead.
func moveZeroes(nums []int) {
	i := 0
	j := 0

	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
}
