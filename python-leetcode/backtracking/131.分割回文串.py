from typing import List


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        res = []
        path = []

        def backtrack(startIndex: int):
            if startIndex == len(s):
                res.append(path[:])
                return
            for i in range(startIndex, len(s)):
                # startIndex - i 是否为回文
                if not self.isPalindrome(s, startIndex, i):
                    continue
                subStr = s[startIndex : i + 1]
                path.append(subStr)
                backtrack(i + 1)
                # 回溯
                path.pop()

        backtrack(0)
        return res

    def isPalindrome(self, s: str, l: int, r: int) -> bool:
        while l < r:
            if s[l] != s[r]:
                return False
            l += 1
            r -= 1
        return True
