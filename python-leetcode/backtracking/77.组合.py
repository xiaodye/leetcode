from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res: List[List[int]] = []
        path: List[int] = []

        def backtrack(start: int):
            if len(path) == k:
                res.append(path[:])
                return

            # 剪枝：最多从 start 到 n - (k - len(path)) + 1
            for i in range(start, n - (k - len(path)) + 2):
                path.append(i)
                backtrack(i + 1)
                path.pop()

        backtrack(1)
        return res
