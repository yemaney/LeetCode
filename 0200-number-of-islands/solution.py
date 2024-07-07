class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        """
        1. traverse the grid using bfs to find continue peices of land
        2. keep track of visited squares of grids in order not to double count
        """

        ans : int = 0
        rows, cols = len(grid), len(grid[0])
        visited :set[tuple[int]] = set()


        def bfs(row: int, col: int):
            q : deque = deque()
            q.append((row, col))
            visited.add((row, col))

            directions : list[tuple[int]] = [(1,0),(0,1),(-1,0),(0,-1)]  
            
            while q:
                row, col = q.popleft()
                for dx, dy in directions:
                    r, c = row + dx, col + dy

                    if (0 <= r < rows) and (0 <= c < cols) and (grid[r][c] == "1") and ((r, c) not in visited):
                        q.append((r, c))
                        visited.add((r, c))



        for r in range(rows):
            for c in range(cols):
                if (grid[r][c] == "1") and ((r, c) not in visited):
                    ans += 1
                    bfs(r, c)


        return ans
