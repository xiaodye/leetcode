from typing import List


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        i = 0
        j = 0

        # 第一趟：把 0 换到前面
        while j < len(nums):
            if nums[j] == 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j += 1
            else:
                j += 1

        # 第二趟：把 1 换到 0 的后面
        j = i
        while j < len(nums):
            if nums[j] == 1:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j += 1
            else:
                j += 1
