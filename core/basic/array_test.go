package basic

import (
	"go-learn/util"
	"testing"
)

func TestArray(t *testing.T) {
	// intArray 类型：[3]int
	intArray := [...]int{1, 2, 3}
	intSlice := intArray[:]

	// 数组不同于切片的很重要的一点，数组为值类型 - 会发生拷贝
	intArray2 := intArray
	util.PrintSliceInfo(intArray)
	util.PrintSliceInfo(intArray2)

	func(intArray [3]int) {
		util.PrintSliceInfo(intArray)
	}(intArray)

	_ = intSlice
}
