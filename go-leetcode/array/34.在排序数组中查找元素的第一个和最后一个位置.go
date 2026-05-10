package leetcode

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			i, j := 1, 1

			// 向左扩展
			for mid >= i && nums[mid-i] == target {
				i++
			}
			// 向右扩展（增加边界判断）
			for mid+j < len(nums) && nums[mid+j] == target {
				j++
			}
			return []int{mid - i + 1, mid + j - 1}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}
