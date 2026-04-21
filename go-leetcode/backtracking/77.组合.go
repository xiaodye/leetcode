package leetcode

func combine(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			// 深拷贝路径
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		// 剪枝：最多从 start 到 n-(k-len(path))+1
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)
	return result
}
