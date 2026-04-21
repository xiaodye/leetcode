package leetcode

// knapsack 0-1背包问题
// weight: 物品重量数组
// value:  物品价值数组
// c:      背包容量
// 返回最大总价值
func knapsack(weight []int, value []int, c int) int {
	n := len(weight)
	if n == 0 {
		return 0
	}

	// 创建二维 dp 数组，全部初始化为 0
	dp := make([][]int, n)

	for i := range n {
		dp[i] = make([]int, c+1)
	}

	// 初始化第一行
	for j := weight[0]; j <= c; j++ {
		dp[0][j] = value[0]
	}

	// 动态规划
	for i := 1; i < n; i++ {
		for j := 0; j <= c; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 不选 i 或者选 i
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[n-1][c]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
