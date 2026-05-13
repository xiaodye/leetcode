package leetcode

func generate(numRows int) [][]int {
	res := [][]int{}

	for i := 0; i < numRows; i++ {
		// 每行 + 1  个数
		rows := make([]int, i+1)
		for idx := range rows {
			rows[idx] = 1
		}

		for j := 1; j < len(rows)-1; j++ {
			rows[j] = res[i-1][j-1] + res[i-1][j]
		}

		res = append(res, rows)
	}

	return res
}
