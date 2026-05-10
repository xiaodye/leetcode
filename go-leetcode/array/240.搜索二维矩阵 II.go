package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	// 对每一层进行二分查找
	// 时间复杂度 O(mlogn)
	for _, nums := range matrix {
		idx := search(nums, target)
		if idx != -1 {
			return true
		}
	}
	return false
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
