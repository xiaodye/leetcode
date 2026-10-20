class Solution:
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        # 单调递减栈，栈存储索引
        stack: list[int] = []

        # 初始化 0
        res = [0] * len(temperatures)

        for i, t in enumerate(temperatures):
            # 维护一个单调递减栈
            while stack and temperatures[stack[-1]] < t:
                top = stack.pop()
                res[top] = i - top

            # 进栈
            stack.append(i)

        return res
