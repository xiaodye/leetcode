package leetcode

func findDuplicate(nums []int) int {
	set := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if !set[nums[i]] {
			set[nums[i]] = true
		} else {
			return nums[i]
		}
	}
	return 0
}
