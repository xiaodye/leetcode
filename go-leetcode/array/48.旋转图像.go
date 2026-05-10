package leetcode

func rotateMatrix(matrix [][]int) {
	// 主要看旋转后的坐标变化 matrix[i][j] -> matrix[j][len - 1 - i]
	// 直接换会导致有些数据会被替换，需要创建新数组
	n := len(matrix)
	newMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		newMatrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newMatrix[j][n-1-i] = matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = newMatrix[i][j]
		}
	}
}
