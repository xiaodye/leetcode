from typing import List


class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res: List[List[int]] = []
        path: List[int] = []
        used = [False] * len(nums)

        def backtrack(start: int):
            # 将当前子集的副本加入结果
            res.append(path[:])

            for i in range(start, len(nums)):
                # 树层去重：如果当前元素与前一个相同且前一个未被使用，则跳过
                if i > 0 and nums[i] == nums[i - 1] and not used[i - 1]:
                    continue

                path.append(nums[i])
                used[i] = True
                backtrack(i + 1)
                path.pop()
                used[i] = False

        backtrack(0)
        return res
