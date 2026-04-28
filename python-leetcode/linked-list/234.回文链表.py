from typing import Optional

from type import ListNode


class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        # 1. 将链表值存入列表
        values: list[int] = []
        curr = head

        while curr:
            values.append(curr.val)
            curr = curr.next

        # 2. 双指针检查回文
        i, j = 0, len(values) - 1
        while i < j:
            if values[i] != values[j]:
                return False
            i += 1
            j -= 1
        return True
