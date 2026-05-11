from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []

        def backtrack(left: int, right: int, s: str):
            # 边界条件
            #  1. 左括号数小于 右，直接结束
            #  2. 括号数量等于 n, push 结果，结束
            if left < right:
                return
            if left + right == 2 * n:
                res.append(s)
                return
            if left < n:
                backtrack(left + 1, right, s + "(")
            if right < n:
                backtrack(left, right + 1, s + ")")

        backtrack(0, 0, "")
        return res
