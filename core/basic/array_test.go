package basic

import (
	"go-learn/tool"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	grammeCandyTest(t)

	arrayIsValueTest(t)
	// arrayToSliceTest(t)
	// sliceToArrayTest(t)
}

func grammeCandyTest(t *testing.T) {
	// 数组语法糖
	var _ = [...]int{1} // [1]int
	var _ = [2]int{1}   // [2]int
}

// 数组不同于切片的很重要的一点，数组为值类型 - 会发生拷贝
func arrayIsValueTest(t *testing.T) {
	// 引用类型
	slice := []int{1, 2, 3}
	tool.PrintAddr(slice)

	// 值类型（这个打印结果，希望你明白，是类型不匹配的意思）
	array := [3]int{1, 2, 3}
	tool.PrintAddr(array)
}

func arrayToSliceTest(t *testing.T) {
	array := [...]int{1, 2, 3}
	slice := array[:]

	t.Log(reflect.TypeOf(slice).String())
}

// Go 1.17 后支持将 slice 转成 array 了（记得修改 go.mod 中定义的版本）
func sliceToArrayTest(t *testing.T) {
	slice := []int{1, 2, 3}
	array := *(*[3]int)(slice)

	t.Log(reflect.TypeOf(array).String())
}
