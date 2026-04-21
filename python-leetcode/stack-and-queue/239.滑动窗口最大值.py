from collections import deque
from typing import List


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        res: List[int] = []
        queue = deque()  # 存储索引

        for i in range(len(nums)):
            # 保持队列单调递减：移除队尾所有小于当前值的索引
            while queue and nums[i] > nums[queue[-1]]:
                queue.pop()
            queue.append(i)

            # 移除窗口外的索引
            if queue[0] <= i - k:
                queue.popleft()

            # 当窗口形成时，记录最大值
            if i >= k - 1:
                res.append(nums[queue[0]])

        return res
