from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        directions = [
            [0, 1],  # 右
            [1, 0],  # 下
            [0, -1],  # 左
            [-1, 0],  # 上
        ]

        count = 0
        rows = len(grid)
        cols = len(grid[0])

        def dfs(i: int, j: int) -> None:
            if i < 0 or i >= rows or j < 0 or j >= cols or grid[i][j] == "0":
                return
            grid[i][j] = "0"
            for dx, dy in directions:
                dfs(i + dx, j + dy)

        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == "1":
                    dfs(i, j)
                    count += 1
        return count
