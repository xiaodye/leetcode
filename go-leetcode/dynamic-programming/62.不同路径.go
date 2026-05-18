package leetcode

func uniquePaths(m int, n int) int {
	// dp[i][j] 机器人到 （i,j）的路径数量
	// dp[i][j] = dp[i -1][j] + dp[i][j-1]
	// 初始化，第一行第一列都为 1
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := range m {
		dp[i][0] = 1
	}

	for j := range n {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
