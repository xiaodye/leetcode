from typing import Optional

from type import ListNode


class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None

        slow = fast = head

        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                # 相遇，找环入口
                x = head
                y = slow
                while x != y:
                    assert x is not None and y is not None  # 保证不为 None
                    x = x.next
                    y = y.next
                return x
        return None
