package _2unique_paths

// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.
//
// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
//
// The test cases are generated so that the answer will be less than or equal to 2 * 10^9.
//
//
// Constraints:
//
//    1 <= m, n <= 100

// [问题]
// m 是纵向的宽、n 是横向的长，左下角想要到达右上角独一无二的路径数量有多少
func uniquePaths(m int, n int) int {
	// return fail01(m, n)
	return step02(m, n)
}

// 排列组合计算表达式的分母最后肯定能够全部约分掉，不然结果是不会为整数的
//  → 但是实现约分也很有难度，难道说这道题不能用排列组合的规律去做？
//  → 发现了一个方格规律，实现的话，时间复杂度也变成了 O(m*n)
// [参考]
// - 参考最优时间复杂度的做法，发现初始化每行的空间，直接在开头启了一个循环去做，本质做法和当前一样
// - 最佳的空间复杂度的做法，果然通过了一种巧妙的方式，将下面自己 fail01 的做法做出来了
func step02(m, n int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				row[j] = 1
			} else {
				row[j] = row[j-1] + arr[i-1][j]
			}
		}
		arr[i] = row
	}
	return arr[n-1][m-1]
}

// 好像就是一个排列组合的数学题：C (m-1) (m+n-2)
// A_m+n_m+n / (A_m_m * A_n_n) 的做法会导致计算过程中分子的数量过大
// A_m+n_n / A_n_n 也有同样的问题
func fail01(m int, n int) int {
	a, b := 1, 1
	s := m + n - 2
	for i := 1; i <= m-1; i++ {
		a *= s
		s--

		b *= i
	}
	return a / b
}
