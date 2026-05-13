package leetcode

func rob(nums []int) int {
	// dp 的含义 dp[i] 盗取 0- i 间房的最大金额
	// dp 递推公式 dp[i] = max(dp[i - 1], dp[i - 2] + nums[i]) 盗取 i 或者不盗取
	// dp 初始值 dp[0] = nums[0], dp[1] = max(nums[0], nums[1])
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}
