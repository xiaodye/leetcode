package leetcode

func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// 将当前子集的副本加入结果
		temp := make([]int, len(subset))
		copy(temp, subset)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			subset = append(subset, nums[i])
			backtrack(i + 1)
			subset = subset[:len(subset)-1] // 回溯
		}
	}

	backtrack(0)
	return res
}
