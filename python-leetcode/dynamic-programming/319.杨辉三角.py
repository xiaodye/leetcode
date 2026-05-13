from typing import List


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        res = []

        for i in range(numRows):
            # 每行 + 1  个数
            rows = [1] * (i + 1)

            for j in range(1, len(rows) - 1):
                rows[j] = res[i - 1][j - 1] + res[i - 1][j]

            res.append(rows)

        return res
