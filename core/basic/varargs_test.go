package basic

import (
	"testing"
)

/*
没有可变参数类型（unpackArray := array...），这只是函数参数的一种语法糖
*/

func printVarargs(array ...int) {
	println(len(array), cap(array))
}

func printSlice(array []int) {
	println(len(array), cap(array))
}

func TestVarargs(t *testing.T) {
	sliceTest()
	arrayTest()
	compareTest()
}

func bestPractice() {
	bad := func(def string, args ...string) string {
		var param string
		if len(args) == 0 {
			param = def
		} else {
			param = args[0]
		}
		return param
	}
	good1 := func(def string, args ...string) string {
		return append(args, def)[0]
	}
	good2 := func(args ...string) string {
		def := "123"
		return append(args, def)[0]
	}

	_, _, _ = bad, good1, good2
}

func sliceTest() {
	var slice = []int{1, 2, 3}
	printVarargs(1, 2, 3)
	printVarargs(slice...)

	// 可变形参类型的形参，不能直接接收切片类型，需要在参数后边加上 3 个点
	// printVarargs(slice)

	// 加上 3 个点，也不能再添加其他参数了
	// printVarargs(slice..., 4)
	// printVarargs(4, slice...)
	printVarargs(append(slice, 4)...)
}

func arrayTest() {
	var array = [3]int{1, 2, 3}
	printVarargs(array[:]...)
	printSlice(array[:])

	// 数组不支持可变参数（语法糖实锤）
	// printVarargs(array...)
}

// 原生函数的泛型行为对比
func compareTest() {
	f := func(_ ...interface{}) {}

	slice := []int{1, 2, 3}

	// 编译不通过
	// f(slice...)
	// 编译通过
	f(slice)
	// 同理理解，居然编译不通过，这说明 append 的泛型参数类型 Type 和 interface 类型本质上还是不同的
	// slice = append([]interface{}{1, 2, 3}, slice...)
}
