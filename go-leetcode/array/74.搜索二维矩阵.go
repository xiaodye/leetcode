package leetcode

func searchMatrixOrder(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		// 判断是否大于该行最大元素，大于则不会在该行
		if target > matrix[i][n-1] {
			continue
		}

		l := 0
		r := n - 1

		for l <= r {
			mid := l + (r-l)/2
			if matrix[i][mid] > target {
				r = mid - 1
			} else if matrix[i][mid] < target {
				l = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}
