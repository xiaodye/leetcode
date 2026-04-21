package leetcode

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	// 计算总和
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 总和为奇数时不可能平分
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	// 如果最大元素超过 target，不可能
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	if maxNum > target {
		return false
	}

	// 创建 dp 数组，dp[i][j] 表示前 i 个元素能否凑出和为 j
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true // 和为 0 总是成立
	}

	// 初始化第一行
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	// 动态规划
	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target]
}
