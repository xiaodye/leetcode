from typing import List


class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        set_ = set()

        for i in range(len(nums)):
            if nums[i] not in set_:
                set_.add(nums[i])
            else:
                return nums[i]
        return 0
