from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        res = []
        path = []

        def backtrack(startIndex: int, total: int):
            # 大于 target ,结束
            if total > target:
                return
            # 等于，算结果
            if total == target:
                res.append(path[:])
                return
            # 元素可以无限用，所以每次都 i = startIndex
            for i in range(startIndex, len(candidates)):
                path.append(candidates[i])
                backtrack(i, total + candidates[i])
                path.pop()

        backtrack(0, 0)
        return res
