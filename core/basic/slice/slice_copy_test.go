package slice

import (
	"testing"
)

func TestCopy(t *testing.T) {
	source := []int{1, 2, 3}

	var target1 []int
	target2 := make([]int, 0)
	target3 := make([]int, 0, 3)
	target4 := make([]int, 3)

	println(copy(target1, source)) // 0
	println(copy(target2, source)) // 0
	println(copy(target3, source)) // 0
	println(copy(target4, source)) // 3

	// 结论：只有有空间的切片或数组 copy 才能够复制
}
