from typing import Optional

from type import ListNode


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(-1, head)
        slow = fast = dummy

        # fast 先走 n 步
        for _ in range(n):
            assert fast is not None
            fast = fast.next

        # fast 再提前走一步，让 slow 指向待删节点的前驱
        assert fast is not None
        fast = fast.next

        # 同时移动
        while fast:
            assert slow is not None
            slow = slow.next
            fast = fast.next

        # 删除 slow.next
        assert slow is not None
        assert slow.next is not None
        slow.next = slow.next.next

        return dummy.next
