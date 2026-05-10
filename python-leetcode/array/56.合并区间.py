from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        result = []

        # 排序
        intervals.sort(key=lambda x: x[0])

        for i in range(len(intervals)):
            if i == 0:
                result.append(intervals[i])
            elif result[-1][1] < intervals[i][0]:
                result.append(intervals[i])
            else:
                arr = result.pop()
                result.append([arr[0], max(arr[1], intervals[i][1])])

        return result
