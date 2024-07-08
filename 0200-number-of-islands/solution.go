func numIslands(grid [][]byte) int {
	/*
	1. traverse the grid using bfs to find continue peices of land
	2. keep track of visited squares by setting their value to '2'
	*/
	ans := 0
	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				ans += 1
				bfs(grid, r, c)
			}
		}
	}

	return ans
}


func bfs(grid [][]byte, row, col int) {
	q := [][2]int{}
	q = append(q, [2]int{row, col})
	grid[row][col] = '2'

	directions := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(q) > 0 {
		rc := q[0]
		q = q[1:]
		for _, v := range directions {

			nr, nc := rc[0]+v[0], rc[1]+v[1]

			if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) && grid[nr][nc] == '1' {
				q = append(q, [2]int{nr, nc})
				grid[nr][nc] = '2'
			}

		}
	}

}

