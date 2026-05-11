class Solution:
    def decodeString(self, s: str) -> str:
        num_stack = []
        str_stack = []
        num = 0
        res = ""

        for ch in s:
            if ch.isdigit():
                # 外层数字情况，可能大于 10，需要累乘 123[a]
                num = num * 10 + int(ch)
            elif ch == "[":
                str_stack.append(res)
                res = ""
                num_stack.append(num)
                num = 0
            elif ch == "]":
                repeat_times = num_stack.pop()
                prev = str_stack.pop()
                res = prev + res * repeat_times
            else:
                # 里层字符情况，可能存在多个
                res += ch

        return res
