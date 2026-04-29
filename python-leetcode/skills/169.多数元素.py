from typing import List


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        # 遍历数组，记录每个数字出现的次数
        freq = {}
        for num in nums:
            freq[num] = freq.get(num, 0) + 1

        # 遍历map, 找出count > arr.length/2 的数字
        for num, count in freq.items():
            if count > len(nums) // 2:
                return num
        return -1
