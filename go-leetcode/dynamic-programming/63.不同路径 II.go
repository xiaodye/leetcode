package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 1. dp[i][j] 的含义
	// 2. dp[i][j] 递推公式 dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	// 3. 初始化 dp[i][j]
	// 4. 遍历
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一列
	for i := 0; i < m && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	// 初始化第一行
	for j := 0; j < n && obstacleGrid[0][j] != 1; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
