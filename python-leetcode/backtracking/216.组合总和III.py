from typing import List


class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        res: List[List[int]] = []
        path: List[int] = []

        def backtrack(start: int, current_sum: int):
            if len(path) == k:
                if current_sum == n:
                    res.append(path[:])
                return

            # 剪枝：最多到 9 - (k - len(path)) + 1
            for i in range(start, 9 - (k - len(path)) + 2):
                path.append(i)
                backtrack(i + 1, current_sum + i)
                path.pop()

        backtrack(1, 0)
        return res
