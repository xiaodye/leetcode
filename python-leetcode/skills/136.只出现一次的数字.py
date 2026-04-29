from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        """
        找出数组中只出现一次的数字（其他数字均出现两次）
        利用异或运算：a ^ a = 0，a ^ 0 = a，且满足交换律与结合律
        """
        res = 0
        for num in nums:
            res ^= num
        return res
