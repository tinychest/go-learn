package _00number_of_islands

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
//
//
// Constraints:
//
//    m == grid.length
//    n == grid[i].length
//    1 <= m, n <= 300
//    grid[i][j] is '0' or '1'.

// - 思路1：遍历所有方格，如果是没有遍历过的方格就进行拓展遍历
func numIslands(grid [][]byte) int {
	res := 0

	x, y := len(grid), len(grid[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == '1' {
				res += 1
				// 记录访问过的节点
				grid[i][j] = '0'
				// 将相关的陆地方块都置为 0
				expand(i, j, grid)
			}
		}
	}
	return res
}

func expand(i, j int, grid [][]byte) {
	x, y := len(grid), len(grid[0])

	grid[i][j] = '0'

	if i > 0 && grid[i-1][j] == '1' {
		expand(i-1, j, grid)
	}
	if i < x-1 && grid[i+1][j] == '1' {
		expand(i+1, j, grid)
	}
	if j > 0 && grid[i][j-1] == '1' {
		expand(i, j-1, grid)
	}
	if j < y-1 && grid[i][j+1] == '1' {
		expand(i, j+1, grid)
	}
}
