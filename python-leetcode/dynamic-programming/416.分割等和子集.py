from typing import List


class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        n = len(nums)
        if n < 2:
            return False

        total = sum(nums)
        # 总和为奇数无法平分
        if total % 2 != 0:
            return False
        target = total // 2

        # 最大元素超过目标值则无法平分
        if max(nums) > target:
            return False

        # 创建 dp 表，dp[i][j] 表示前 i 个元素能否凑出和为 j
        dp = [[False] * (target + 1) for _ in range(n)]
        for i in range(n):
            dp[i][0] = True  # 和为 0 总是可以做到

        # 初始化第一行
        if nums[0] <= target:
            dp[0][nums[0]] = True

        # 填充 dp 表
        for i in range(1, n):
            for j in range(1, target + 1):
                if nums[i] > j:
                    dp[i][j] = dp[i - 1][j]
                else:
                    dp[i][j] = dp[i - 1][j] or dp[i - 1][j - nums[i]]

        return dp[n - 1][target]
