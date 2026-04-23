from typing import List


class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        # left组合 + right组合 = sum
        # 需要 left组合 - right组合 = target
        # 推导出 left组合 = (target + sum) / 2， 0 - 1 背包问题
        total = sum(nums)

        if total < abs(target):
            return 0
        if (target + total) % 2 == 1:
            return 0

        # x 可以看成是背包容量，需要从 nums 中凑出 x，01背包问题
        # dp[i][j]: 前 i 个物品是否能凑出 j
        x = (target + total) // 2

        n = len(nums)
        # 创建二维 dp 数组
        dp = [[0] * (x + 1) for _ in range(n)]

        # 初始化dp, dp[0][0] = 1, dp[0][nums[0]] = 1;
        if nums[0] <= x:
            dp[0][nums[0]] = 1
        dp[0][0] = 1

        if nums[0] == 0:
            dp[0][0] += 1

        # 第2行和第1列开始
        for i in range(1, n):
            for j in range(x + 1):
                if j >= nums[i]:
                    dp[i][j] = dp[i - 1][j] + dp[i - 1][j - nums[i]]
                else:
                    dp[i][j] = dp[i - 1][j]

        return dp[n - 1][x]
