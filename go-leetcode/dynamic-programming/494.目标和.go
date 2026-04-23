package leetcode

func findTargetSumWays(nums []int, target int) int {
	// left组合 + right组合 = sum
	// 需要 left组合 - right组合 = target
	// 推导出 left组合 = (target + sum) / 2， 0 - 1 背包问题
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum < abs(target) {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}

	// x 可以看成是背包容量，需要从 nums 中凑出 x，01背包问题
	// dp[i][j]: 前 i 个物品是否能凑出 j
	x := (target + sum) / 2

	n := len(nums)
	// 创建二维 dp 数组，类型为 int
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, x+1)
	}

	// 初始化dp, dp[0][0] = 1, dp[0][nums[0]] = 1;
	if nums[0] <= x {
		dp[0][nums[0]] = 1
	}
	dp[0][0] = 1

	if nums[0] == 0 {
		dp[0][0]++
	}

	// 第2行和第1列开始
	for i := 1; i < n; i++ {
		for j := 0; j <= x; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n-1][x]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
