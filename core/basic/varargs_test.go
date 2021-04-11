package basic

import (
	"testing"
)

func varargsBasicConcept() {
	// 没有可变参数类型
	// unpackArray := array...
}

func printVarargs(array ...int) {
	println(len(array), cap(array))
}

func printSlice(array []int) {
	println(len(array), cap(array))
}

func TestVarargs(t *testing.T) {
	var slice = []int{1, 2, 3}
	var array = [3]int{1, 2, 3}

	// 可变形参类型的形参，不能直接接收切片，需要在参数后边加上 3 个点
	// printVarargs(slice)
	// 数组直接凉凉，加点也不行
	// printVarargs(array...)
	// 但是数组转化成切片再进行相关的操作
	printVarargs(array[:]...)
	printSlice(array[:])

	printVarargs(1, 2, 3)

	printVarargs(slice...)

	printSlice(slice)
}

func TestVarargs2(t *testing.T) {
	theFunc := func(interfaces ...interface{}) {
		// interfaces 参数后边不加 ...，得到提示：missing ... in args forwarded to print-like function
		println(interfaces)
	}

	var array = []interface{}{1, 2, 3}
	theFunc(array)
}

func TestVarargs3(t *testing.T) {
	varargsParamFunc := func(_ ...interface{}) {}

	intSlice := []int{1, 2, 3}

	// 编译通过
	varargsParamFunc(intSlice)
	// 同理理解，居然编译不通过，这说明 append 的泛型参数类型 Type 和 interface 类型本质上还是不同的
	// intSlice = append([]interface{}{1, 2, 3}, intSlice...)
}
