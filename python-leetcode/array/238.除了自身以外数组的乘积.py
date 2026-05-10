from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        # 前缀和
        # prefix[i]: i 左侧的数组乘积和
        # postfix[i]: i 右侧的数组的乘积和
        # answer[i] = prefix[i] * postfix[i]
        prefix = [0] * len(nums)
        postfix = [0] * len(nums)
        answer = [0] * len(nums)

        prefix[0] = 1
        postfix[len(nums) - 1] = 1

        for i in range(1, len(nums)):
            prefix[i] = prefix[i - 1] * nums[i - 1]

        for i in range(len(nums) - 2, -1, -1):
            postfix[i] = postfix[i + 1] * nums[i + 1]

        for i in range(len(nums)):
            answer[i] = prefix[i] * postfix[i]

        return answer
