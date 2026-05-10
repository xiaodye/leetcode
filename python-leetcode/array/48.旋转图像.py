from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        # 主要看旋转后的坐标变化 matrix[i][j] -> matrix[j][len - 1 - i]
        # 直接换会导致有些数据会被替换，需要创建新数组
        n = len(matrix)
        new_matrix = [[0] * n for _ in range(n)]

        for i in range(n):
            for j in range(n):
                new_matrix[j][n - 1 - i] = matrix[i][j]

        for i in range(n):
            for j in range(n):
                matrix[i][j] = new_matrix[i][j]
