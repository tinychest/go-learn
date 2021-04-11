package slice

import (
	"testing"
)

func TestCopy(t *testing.T) {
	sourceIntSlice := []int{1, 2, 3}

	var targetIntSlice1 []int
	targetIntSlice2 := make([]int, 0)
	targetIntSlice3 := make([]int, 0, 3)
	targetIntSlice4 := make([]int, 3)

	println(copy(targetIntSlice1, sourceIntSlice)) // 0
	println(copy(targetIntSlice2, sourceIntSlice)) // 0
	println(copy(targetIntSlice3, sourceIntSlice)) // 0
	println(copy(targetIntSlice4, sourceIntSlice)) // 3

	// 结论：只有有空间的切片或数组 copy 才能够复制
}
