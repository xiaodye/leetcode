from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        # 问题其实为 取一个数左边，一个数右边，最大差值
        # 找最小值一定是必须且固定的
        min_price = prices[0]  # 初始化为第一个价格，避免 float('inf') 导致的类型警告
        res = 0

        for i in range(len(prices)):
            min_price = min(min_price, prices[i])
            res = max(res, prices[i] - min_price)

        return res
