from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []

        phone_map: dict[str, str] = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }

        res: List[str] = []

        def backtrack(index: int, path: str):
            if len(path) == len(digits):
                res.append(path)
                return

            letters = phone_map[digits[index]]
            for char in letters:
                backtrack(index + 1, path + char)

        backtrack(0, "")
        return res
