package _8rotate_image

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
//
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//
//
// Constraints:
//
//    n == matrix.length == matrix[i].length
//    1 <= n <= 20
//    -1000 <= matrix[i][j] <= 1000

// 整体的旋转就是，每个数环旋转，那么其实只要提供一个 四个数交换值的方法就行了
// 花了超多时间来处理细节：
// - 参数二维数组的第一个中括号实际上是纵坐标，第二个中括号才是横坐标
// - 在环层向内缩变化的情况下，还要知道准确的坐标，很容易就犯错；自己还为了简化引入概念，结果混杂在一起
// [结果]
// 第一次提交 4ms，第二次提交 0ms ...
// 看不了参考答案，请求显示 500 了
func rotate(matrix [][]int) {
	l := len(matrix)

	// 有几层需要旋转
	for i := 0; i < l/2; i++ {
		// 每层需要进行 4 旋的方块
		for j := i; j < l-i-1; j++ {
			// 四旋
			t := matrix[j][i] // 左下角提起来 ∟

			matrix[j][i] = matrix[l-i-1][j]         // [右]下到左[下]
			matrix[l-i-1][j] = matrix[l-j-1][l-i-1] // 右[上]到[右]下
			matrix[l-j-1][l-i-1] = matrix[i][l-j-1] // [左]上到右[上]
			matrix[i][l-j-1] = t                    // 左[下]到[左]上
			// fmt.Printf("%d,%d [%d] = %d,%d [%d]\n", j, i, t, le, j, matrix[le][j])
			// fmt.Printf("%d,%d [%d] = %d,%d [%d]\n", le, j, matrix[le][j], le-j, le-i, matrix[le-j][le-i])
			// fmt.Printf("%d,%d [%d] = %d,%d [%d]\n", le-j, le-i, matrix[le-j][le-i], i, le-j, matrix[i][le-j])
			// fmt.Printf("%d,%d [%d] = %d,%d [%d]\n", i, le-j, matrix[i][le-j], j, i, t)
		}
	}
}
