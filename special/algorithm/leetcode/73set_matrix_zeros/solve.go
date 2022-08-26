package _3set_matrix_zeros

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
//
// You must do it in place.
//
//
// Constraints:
//
//    m == matrix.length
//    n == matrix[0].length
//    1 <= m, n <= 200
//    -2^31 <= matrix[i][j] <= 2^31 - 1

// 这个题没什么内容，一个点是，map 的值使用 struct 性能比 bool 差一些
func setZeroes(matrix [][]int) {
	xm := make(map[int]bool)
	ym := make(map[int]bool)

	xv, yv := len(matrix), len(matrix[0])

	for i := 0; i < xv; i++ {
		for j := 0; j < yv; j++ {
			if matrix[i][j] == 0 {
				xm[i] = true
				ym[j] = true
			}
		}
	}

	for i := 0; i < xv; i++ {
		for j := 0; j < yv; j++ {
			if xm[i] || ym[j] {
				matrix[i][j] = 0
			}
		}
	}
}
