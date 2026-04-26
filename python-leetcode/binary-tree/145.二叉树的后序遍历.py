from typing import List, Optional

from type import TreeNode


class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res = []

        def postorder(node: Optional[TreeNode]) -> None:
            if not node:
                return
            postorder(node.left)
            postorder(node.right)
            res.append(node.val)

        postorder(root)
        return res
