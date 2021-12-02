package basic

import (
	"fmt"
	"go-learn/util"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	arrayIsValueTest()
	// arrayToSliceTest()
	// sliceToArrayTest()
}

// 数组不同于切片的很重要的一点，数组为值类型 - 会发生拷贝
func arrayIsValueTest() {
	// 引用类型
	slice := []int{1, 2, 3}
	util.PrintAddr(slice)

	// 值类型（这个打印结果，希望你明白，是类型不匹配的意思）
	array := [3]int{1, 2, 3}
	util.PrintAddr(array)
}

func arrayToSliceTest() {
	array := [...]int{1, 2, 3}
	slice := array[:]

	fmt.Println(reflect.TypeOf(slice).String())
}

// Go 1.17 后支持将 slice 转成 array 了（记得修改 go.mod 中定义的版本）
func sliceToArrayTest() {
	slice := []int{1, 2, 3}
	array := *(*[3]int)(slice)

	fmt.Println(reflect.TypeOf(array).String())
}
