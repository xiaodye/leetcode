from typing import List


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        # dp[i] 表示金额之和等于 i 的硬币组合数
        dp = [0] * (amount + 1)
        dp[0] = 1  # 金额为0，只有一种方式：不放任何硬币

        for coin in coins:
            for j in range(1, amount + 1):
                if coin <= j:
                    dp[j] += dp[j - coin]
        return dp[amount]
