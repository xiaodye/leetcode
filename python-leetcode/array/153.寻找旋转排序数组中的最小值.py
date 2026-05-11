from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        # 旋转后，还是原数组情况，元素升序
        if len(nums) == 1 or nums[0] < nums[-1]:
            return nums[0]

        left, right = 0, len(nums) - 1

        while left < right:
            mid = left + (right - left) // 2

            # 前半段无序，忽略后半段
            if nums[0] <= nums[mid]:
                left = mid + 1
            else:
                right = mid

        return nums[left]
