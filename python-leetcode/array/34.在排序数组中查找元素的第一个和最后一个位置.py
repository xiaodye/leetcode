from typing import List


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        left, right = 0, len(nums) - 1

        while left <= right:
            mid = left + (right - left) // 2

            if nums[mid] == target:
                i, j = 1, 1

                # 向左扩展
                while mid >= i and nums[mid - i] == target:
                    i += 1
                # 向右扩展
                while mid + j < len(nums) and nums[mid + j] == target:
                    j += 1

                return [mid - i + 1, mid + j - 1]

            elif nums[mid] > target:
                right = mid - 1
            else:
                left = mid + 1

        return [-1, -1]
