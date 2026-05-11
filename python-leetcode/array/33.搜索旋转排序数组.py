from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if len(nums) == 1:
            return 0 if nums[0] == target else -1
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            elif nums[0] <= nums[mid]:
                # 前半段是有序的
                if target >= nums[0] and target < nums[mid]:
                    right = mid - 1
                else:
                    left = mid + 1
            else:
                # 否则，肯定是后半段有序
                if target <= nums[right] and target > nums[mid]:
                    left = mid + 1
                else:
                    right = mid - 1
        return -1
