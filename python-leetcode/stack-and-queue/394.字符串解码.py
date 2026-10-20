class Solution:
    def decodeString(self, s: str) -> str:
        num_stack: list[int] = []
        str_stack: list[str] = []

        curr_num = 0
        curr_str = ""

        for char in s:
            if char.isdigit():
                # 外层数字情况，可能大于 10，需要累乘
                curr_num = curr_num * 10 + int(char)
            elif char == "[":
                # 进栈
                num_stack.append(curr_num)
                str_stack.append(curr_str)
                curr_num = 0
                curr_str = ""
            elif char == "]":
                # 出栈
                count = num_stack.pop()
                prev = str_stack.pop()
                curr_str = prev + curr_str * count
            else:
                # 字母，累加
                curr_str += char

        return curr_str
