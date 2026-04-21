from typing import List


class Solution:
    def fib(self, n: int) -> int:
        if n <= 1:
            return n

        dp: List[int] = [0] * (n + 1)
        dp[1] = 1

        for i in range(2, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]

        return dp[n]
