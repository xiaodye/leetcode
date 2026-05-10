package leetcode

func maxSubArray(nums []int) int {
	// dp[i]: 第 i 项结尾最大子数组和;
	// 递推式: dp[i] = Math.max(dp[i - 1] + nums[i], nums[i])
	// 初始化dp: dp[0] = nums[0];
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}

	return maxSum
}
