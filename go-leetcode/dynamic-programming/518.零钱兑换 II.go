package leetcode

func change(amount int, coins []int) int {
	// dp[i] 表示金额之和等于 i 的硬币组合数
	dp := make([]int, amount+1)
	dp[0] = 1 // 金额为0，只有一种方式：不放任何硬币

	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if coins[i] <= j {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
