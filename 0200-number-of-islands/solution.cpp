class Solution
{
public:
    /**
     *  1. traverse the grid using dfs to find continuous pieces of land
     *  2. mark squares that are already visited by updating their value to '2'
     */
    int numIslands(vector<vector<char>> &grid)
    {
        int ans = 0;

        for (size_t r = 0; r < grid.size(); r++)
        {
            for (size_t c = 0; c < grid[0].size(); c++)
            {
                if (grid[r][c] == '1')
                {
                    ans += 1;
                    dfs(grid, r, c);
                }
            }
        }

        return ans;
    }

private:
    void dfs(vector<vector<char>> &grid, int row, int col)
    {
        if (row < 0 || row >= grid.size() || col < 0 || col >= grid[0].size() || grid[row][col] != '1')
        {
            return;
        }

        grid[row][col] = '2';

        dfs(grid, row + 1, col);
        dfs(grid, row - 1, col);
        dfs(grid, row, col + 1);
        dfs(grid, row, col - 1);
    }
};
