package leetcode

// Do not return anything, modify matrix in-place instead.
func setZeroes(matrix [][]int) {
	// 第一次遍历，发现 matrix[i][j] === 0 ,则设置 rows[i] = true,col[j] = true
	// 第二次遍历，发现存在rows[i] = true ｜ col[j] = true，则置为 0
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}
