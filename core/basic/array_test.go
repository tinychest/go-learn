package basic

import (
	"go-learn/util"
	"testing"
)

func TestArray(t *testing.T) {
	// array 类型：[3]int
	array := [...]int{1, 2, 3}
	slice := array[:]

	// 数组不同于切片的很重要的一点，数组为值类型 - 会发生拷贝
	array2 := array
	util.PrintSlice(array)
	util.PrintSlice(array2)

	func(arr [3]int) {
		util.PrintSlice(arr)
	}(array)

	_ = slice
}
