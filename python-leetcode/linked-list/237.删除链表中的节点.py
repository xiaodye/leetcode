from typing import Optional

from type import ListNode


class Solution:
    def deleteNode(self, node: Optional[ListNode]):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """

        assert node is not None  # 或 if not node: return
        assert node.next is not None
        node.val = node.next.val
        node.next = node.next.next
