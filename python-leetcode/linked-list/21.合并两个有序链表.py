from typing import Optional

from type import ListNode


class Solution:
    def mergeTwoLists(
        self, list1: Optional[ListNode], list2: Optional[ListNode]
    ) -> Optional[ListNode]:
        # 边界情况：两个链表都为空
        if not list1 and not list2:
            return None

        dummy = ListNode(-1)
        curr = dummy

        while list1 and list2:
            if list1.val <= list2.val:
                curr.next = list1
                list1 = list1.next
            else:
                curr.next = list2
                list2 = list2.next
            curr = curr.next

        # 连接剩余部分
        curr.next = list1 or list2

        return dummy.next
