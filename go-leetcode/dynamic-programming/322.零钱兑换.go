package leetcode

func coinChange(coins []int, amount int) int {
	// dp[i]: 总额 i 对应的最少硬币个数
	// 倒推：假设现在已经有了 11 美分，确定有几个硬币
	// 递推式：dp[i] = Math.min(dp[i - amount[0]],dp[i-amount[1]], .....) + 1
	// 初始化 dp, dp[0] = 0;
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // 初始化为不可能的大数
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
