package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}
	used := make([]bool, len(nums))

	sort.Ints(nums) // 排序以便去重

	var backtrack func(start int)
	backtrack = func(start int) {
		// 将当前子集的副本加入结果
		temp := make([]int, len(subset))
		copy(temp, subset)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			// 树层去重：如果当前元素与前一个相同且前一个未被使用，则跳过
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			subset = append(subset, nums[i])
			used[i] = true
			backtrack(i + 1)
			subset = subset[:len(subset)-1]
			used[i] = false
		}
	}

	backtrack(0)
	return res
}
