package leetcode

func numIslands(grid [][]byte) int {
	directions := [][]int{
		{0, 1},  // 右
		{1, 0},  // 下
		{0, -1}, // 左
		{-1, 0}, // 上
	}

	// 初始化岛屿数量，缓存二维数组的行数与列数
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' {
			return
		}
		// 置为 '0'，代表访问过
		grid[i][j] = '0'
		// 向四个方向 dfs
		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1])
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				// 每完成一层 dfs，count 加 1
				count++
			}
		}
	}
	return count
}
