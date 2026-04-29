from typing import List, Optional

from type import TreeNode


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        def buildBST(left: int, right: int) -> Optional[TreeNode]:
            if left > right:
                return None
            mid = left + (right - left) // 2
            root = TreeNode(nums[mid])
            root.left = buildBST(left, mid - 1)
            root.right = buildBST(mid + 1, right)
            return root

        return buildBST(0, len(nums) - 1) if nums else None
