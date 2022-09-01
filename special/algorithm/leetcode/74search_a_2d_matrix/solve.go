package _4search_a_2d_matrix

// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
//
//    Integers in each row are sorted from left to right.
//    The first integer of each row is greater than the last integer of the previous row.
//
//
// Constraints:
//
//    m == matrix.length
//    n == matrix[i].length
//    1 <= m, n <= 100
//    -104 <= matrix[i][j], target <= 104

// 感觉思路很简单，就是两次二分查找
// [结果] 提交的结果不佳（事后发现其实已经是不错的解法了）
// [参考] 将两次折半结合起来了，实际上虽然是二维的有序数组，但完全可以当作有序的一维数组去看
func searchMatrix(matrix [][]int, target int) bool {
	// return step01(matrix, target)
	return refer(matrix, target)
}

func step01(matrix [][]int, target int) bool {
	start, end := 0, len(matrix)-1
	mid := (start + end) / 2

	// 纵向折半
	for start <= end {
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (start + end) / 2
	}
	if target < matrix[mid][0] {
		return false
	}

	i := mid

	// 横向折半
	start, end = 0, len(matrix[i])-1
	mid = (start + end) / 2

	for start <= end {
		if matrix[i][mid] == target {
			return true
		} else if matrix[i][mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (start + end) / 2
	}
	return false
}

// 提交结果不稳定，一下 9ms，一下 0ms
func refer(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	start, end := 0, rows*cols-1

	for start <= end {
		// 将一维数组概念中的 mid 转化成二位数组中的 横坐标 和 纵坐标
		mid := (start + end) / 2
		row := mid / cols
		col := mid % cols

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
