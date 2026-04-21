from typing import List


class Solution:
    def lastStoneWeightII(self, stones: List[int]) -> int:
        total = sum(stones)
        target = total // 2  # 背包容量

        n = len(stones)
        # 创建 dp 二维列表
        dp = [[0] * (target + 1) for _ in range(n)]

        # 初始化第一行
        if stones[0] <= target:
            for j in range(stones[0], target + 1):
                dp[0][j] = stones[0]

        # 动态规划
        for i in range(1, n):
            for j in range(1, target + 1):
                if j >= stones[i]:
                    dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - stones[i]] + stones[i])
                else:
                    dp[i][j] = dp[i - 1][j]

        return total - 2 * dp[n - 1][target]
