from typing import List


class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        # dp[i][j] 机器人到 （i,j）的路径数量
        # dp[i][j] = dp[i -1][j] + dp[i][j-1]
        # 初始化，第一行第一列都为 1
        dp = [[0] * n for _ in range(m)]
        for i in range(m):
            dp[i][0] = 1
        for j in range(n):
            dp[0][j] = 1

        for i in range(1, m):
            for j in range(1, n):
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        return dp[m - 1][n - 1]
