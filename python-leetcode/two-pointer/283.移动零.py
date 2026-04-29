from typing import List


class Solution:
    """
    Do not return anything, modify nums in-place instead.
    """

    def moveZeroes(self, nums: List[int]) -> None:
        i = 0
        j = 0

        while j < len(nums):
            if nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j += 1
            else:
                j += 1
