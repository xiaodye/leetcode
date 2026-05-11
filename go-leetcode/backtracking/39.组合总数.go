package leetcode

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtrack func(startIndex int, sum int)
	backtrack = func(startIndex int, sum int) {
		// 大于 target ,结束
		if sum > target {
			return
		}
		// 等于，算结果
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		// 元素可以无限用，所以每次都 i = startIndex
		for i := startIndex; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}
