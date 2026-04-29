from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        max_len = 0

        # 把所有节点都存储到 set 中，并进行去重
        s = set(nums)
        nums = list(s)

        # 全部放 set 里
        for i in range(len(nums)):
            # 如果该节点存在前驱节点，说明这个节点开始的序列肯定不是最长的
            if nums[i] - 1 in s:
                continue

            length = 1
            num = nums[i]

            while num + 1 in s:
                length += 1
                num += 1

            max_len = max(max_len, length)

        return max_len
