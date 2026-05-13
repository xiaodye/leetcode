from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        max_jump = 0

        for i in range(len(nums)):
            if i > max_jump:
                return False

            # 我们依次遍历数组中的每一个位置，并实时维护 最远可以到达的位置
            max_jump = max(max_jump, i + nums[i])

            # 最远可以到达的位置 大于等于数组中的最后一个位置
            if max_jump >= len(nums) - 1:
                return True

        return False
