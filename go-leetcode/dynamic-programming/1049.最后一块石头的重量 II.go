package leetcode

func lastStoneWeightII(stones []int) int {
	// 计算总和
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2 // 背包容量

	n := len(stones)
	// 创建 dp 二维数组
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}

	// 初始化第一行
	if stones[0] <= target {
		for j := stones[0]; j <= target; j++ {
			dp[0][j] = stones[0]
		}
	}

	// 动态规划
	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j >= stones[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return sum - 2*dp[n-1][target]
}
