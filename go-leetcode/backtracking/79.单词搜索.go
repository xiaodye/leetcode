package leetcode

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visited[i][j] = true
		res := false
		for _, dir := range directions {
			x := i + dir[0]
			y := j + dir[1]
			if x >= 0 && x < rows && y >= 0 && y < cols && !visited[x][y] {
				if dfs(x, y, k+1) {
					res = true
					break
				}
			}
		}
		visited[i][j] = false
		return res
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
