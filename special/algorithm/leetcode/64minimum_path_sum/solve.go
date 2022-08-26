package _4minimum_path_sum

import "math"

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
//
// Note: You can only move either down or right at any point in time.
//
//
// Constraints:
//
//    m == grid.length
//    n == grid[i].length
//    1 <= m, n <= 200
//    0 <= grid[i][j] <= 100

// 迷宫问题，找到迷宫路径上值和最小的路径
// [思路]
// 迷宫问题肯定是要遍历的，算法的优劣取决于遍历的方式
func minPathSum(grid [][]int) int {
	// return step01(grid)
	// return step02(grid)
	return step03(grid)
}

// 在 step02 的基础上，省去两个循环
// Runtime: 15 ms, faster than 32.12% of Go online submissions for Minimum Path Sum.
// Memory Usage: 3.9 MB, less than 91.82% of Go online submissions for Minimum Path Sum.
// 提升了一点点，不多，到此了，时间复杂度暂时是这样
func step03(grid [][]int) int {
	xv, yv := len(grid), len(grid[0])

	for i := 0; i < xv; i++ {
		for j := 0; j < yv; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[xv-1][yv-1]
}

// [思路]
// 不能简单的递归，效率太低了
// 对于每个方格来说，到达其的最短距离取决于到达 上面 或者 左边 的最短距离
// [结果]
// Runtime: 21 ms, faster than 11.52% of Go online submissions for Minimum Path Sum.
// Memory Usage: 3.9 MB, less than 91.82% of Go online submissions for Minimum Path Sum.
// 参考最佳提交，实际上整合了两个循环
func step02(grid [][]int) int {
	xv, yv := len(grid), len(grid[0])

	for i := 1; i < xv; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < yv; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < xv; i++ {
		for j := 1; j < yv; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[xv-1][yv-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 提交结果：超时
var res int

func step01(grid [][]int) int {
	// 提交 leetcode 中防止测试用例之间互相影响
	res = math.MaxInt
	traverse(grid, 0, 0, 0)
	return res
}

func traverse(grid [][]int, x, y, val int) {
	if x == len(grid)-1 && y == len(grid[0])-1 {
		if res > grid[x][y]+val {
			res = grid[x][y] + val
		}
		return
	}

	// 如果可以往右走，才往右走
	if x < len(grid)-1 {
		traverse(grid, x+1, y, val+grid[x][y])
	}
	// 如果可以往下走，才往下走
	if y < len(grid[0])-1 {
		traverse(grid, x, y+1, val+grid[x][y])
	}
}
