class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) % 2 != 0:
            return False

        pairs = {
            "(": ")",
            "{": "}",
            "[": "]",
        }

        stack = []

        for c in s:
            if c in pairs:
                # 是左括号，进栈
                stack.append(c)
            else:
                # 是右括号，栈空或栈顶不匹配则非法
                if not stack:
                    return False
                if pairs[stack.pop()] != c:
                    return False

        return len(stack) == 0
