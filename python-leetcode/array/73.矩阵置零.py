from typing import List


class Solution:
    """
    Do not return anything, modify matrix in-place instead.
    """

    def setZeroes(self, matrix: List[List[int]]) -> None:
        # 第一次遍历，发现 matrix[i][j] === 0 ,则设置 rows[i] = true,col[j] = true
        # 第二次遍历，发现存在rows[i] = true ｜ col[j] = true，则置为 0
        rows = [False] * len(matrix)
        cols = [False] * len(matrix[0])

        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[i][j] == 0:
                    rows[i] = True
                    cols[j] = True

        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if rows[i] or cols[j]:
                    matrix[i][j] = 0
