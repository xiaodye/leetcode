from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        # dp[i]: 第 i 项结尾最大子数组和;
        # 递推式: dp[i] = Math.max(dp[i - 1] + nums[i], nums[i])
        # 初始化dp: dp[0] = nums[0];
        dp = [0] * len(nums)
        dp[0] = nums[0]
        maxSum = dp[0]

        for i in range(1, len(nums)):
            dp[i] = max(dp[i - 1] + nums[i], nums[i])
            maxSum = max(maxSum, dp[i])

        return maxSum
