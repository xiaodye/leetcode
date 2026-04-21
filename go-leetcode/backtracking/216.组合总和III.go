package leetcode

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtrack func(start int, sum int)
	backtrack = func(start int, sum int) {
		if len(path) == k {
			if sum == n {
				temp := make([]int, k)
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}
		// 剪枝：最多到 9 - (k - len(path)) + 1
		for i := start; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	backtrack(1, 0)
	return res
}
