from typing import Optional

from type import ListNode


class Solution:
    def getIntersectionNode(
        self, headA: ListNode, headB: ListNode
    ) -> Optional[ListNode]:
        visited = set[ListNode]()

        temp = headA
        while temp:
            visited.add(temp)
            temp = temp.next

        temp = headB
        while temp:
            if temp in visited:
                return temp
            temp = temp.next

        return None
