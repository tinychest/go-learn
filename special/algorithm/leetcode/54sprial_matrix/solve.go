package _4sprial_matrix

// 算法题简介：矩阵螺旋遍历
//
// 还可以优化的点：不要一步一步的移动，可以一个方向上的，一段一段的移
//
// 参考答案：比较多的说法是，借助一个

func spiralOrder(matrix [][]int) []int {
	var (
		width                   = len(matrix[0])
		height                  = len(matrix)
		result                  = make([]int, 0, width*height)
		direction directionEnum = RIGHT
		x, y                    = 0, 0
	)
	up, down, left, right = 0, height-1, 0, width-1

	for i := 0; i < width*height; i++ {
		result = append(result, matrix[y][x])
		traverse(&direction, &x, &y)
	}

	return result
}

// 方向枚举
type directionEnum int

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

// 矩阵边界（边界元素下标）
var up, down, left, right int

/* 辅助方法 */
type method interface {
	// 基于当前的坐标和方向移动到下一个正确的坐标
	traverse(direction *directionEnum, x, y *int)
	// 转向，并修改边界
	turn(direction *directionEnum)
	// 以当前的方向和坐标，判断是否可以移动
	canMove(direction *directionEnum, x, y *int) bool
	// 朝着当前方向移动一步
	move(direction *directionEnum, x, y *int)
}

func traverse(direction *directionEnum, x, y *int) {
	for !canMove(direction, x, y) {
		turn(direction)
	}
	move(direction, x, y)
}

func turn(direction *directionEnum) {
	switch *direction {
	case RIGHT:
		up++
	case DOWN:
		right--
	case LEFT:
		down--
	case UP:
		left++
	default:
		panic("never happened")
	}
	*direction = (*direction + 1) % 4
}

func canMove(direction *directionEnum, x, y *int) bool {
	switch *direction {
	case RIGHT:
		return *x != right
	case DOWN:
		return *y != down
	case LEFT:
		return *x != left
	case UP:
		return *y != up
	default:
		panic("never happened")
	}
}

func move(direction *directionEnum, x, y *int) {
	switch *direction {
	case RIGHT:
		*x++
	case DOWN:
		*y++
	case LEFT:
		*x--
	case UP:
		*y--
	}
}
