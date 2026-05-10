from typing import List


class Solution:
    """
    Do not return anything, modify nums in-place instead.
    """

    def rotate(self, nums: List[int], k: int) -> None:
        k = k % len(nums)

        if k == 0:
            return

        # 反转整个数组
        self.reverse(nums, 0, len(nums) - 1)
        # 反转前 k 个
        self.reverse(nums, 0, k - 1)
        # 反转剩余
        self.reverse(nums, k, len(nums) - 1)

    def reverse(self, nums: List[int], start: int, end: int) -> None:
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1
