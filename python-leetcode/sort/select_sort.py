from typing import List

class Solution:
    def selectSort(self, arr: List[int]) -> List[int]:
        if len(arr) == 1:
            return arr

        n = len(arr)

        # 选择排序，从前往后排，外层为轮数
        # 内层为比较次数

        for i in range(n - 1):
            min_idx = i
            for j in range(i + 1, n):
                if arr[j] < arr[min_idx]:
                    min_idx = j
            arr[i], arr[min_idx] = arr[min_idx], arr[i]

        return arr
