from typing import List


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        rows = len(board)
        cols = len(board[0])
        visited = [[False] * cols for _ in range(rows)]
        directions = [
            [0, 1],
            [0, -1],
            [1, 0],
            [-1, 0],
        ]

        def dfs(i: int, j: int, k: int) -> bool:
            if board[i][j] != word[k]:
                return False
            if k == len(word) - 1:
                return True
            visited[i][j] = True
            res = False
            for dx, dy in directions:
                x, y = i + dx, j + dy
                if 0 <= x < rows and 0 <= y < cols and not visited[x][y]:
                    if dfs(x, y, k + 1):
                        res = True
                        break
            visited[i][j] = False
            return res

        for i in range(rows):
            for j in range(cols):
                if dfs(i, j, 0):
                    return True
        return False
