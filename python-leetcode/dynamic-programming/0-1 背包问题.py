from typing import List


def knapsack(weight: List[int], value: List[int], c: int) -> int:
    n = len(weight)
    if n == 0:
        return 0

    # 创建二维 dp 列表，全部初始化为 0
    dp = [[0] * (c + 1) for _ in range(n)]

    # 初始化第一行
    for j in range(weight[0], c + 1):
        dp[0][j] = value[0]

    # 动态规划
    for i in range(1, n):
        for j in range(0, c + 1):
            if weight[i] > j:
                dp[i][j] = dp[i - 1][j]
            else:
                dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i])

    return dp[n - 1][c]
