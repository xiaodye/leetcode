from typing import List


class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        res = [-1] * len(nums1)
        stack: List[int] = []
        mapping = {}

        for i in range(len(nums2)):
            while stack and nums2[i] > nums2[stack[-1]]:
                top_index = stack.pop()
                mapping[nums2[top_index]] = nums2[i]
            stack.append(i)

        for i in range(len(nums1)):
            if nums1[i] in mapping:
                res[i] = mapping[nums1[i]]

        return res
