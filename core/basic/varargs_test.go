package basic

import (
	"testing"
)

// 可变参数类型（unpackArray := array...）不是一个具体存在的类型，这只是函数参数的一种语法糖
// - 不能直接接收切片类型，需要在参数后边加上 ...
// - 顺序上只能放在最后，否则 “Can only use '...' as the final argument in the list”

func TestIntVarargs(t *testing.T) {
	var slice = []int{1, 2, 3}
	var array = [3]int{1, 2, 3}

	f := func(_ ...int) {}

	// 多个元素参数
	f(1, 2, 3)
	// 切片
	f(slice...)
	// 切片 + 多个元素参数
	f(append(slice, 4)...)
	// 数组不支持可变参数，需要先转化成切片
	f(array[:]...)
}

func TestVarargs(t *testing.T) {
	intSlice := []int{1, 2, 3}

	f := func(args ...interface{}) {
		println(len(args))
	}

	// 误区，并不是每个 int 元素赋值到对应的 interface{} 类型元素；而是 intSlice 整个作为一个 interface{} 类型的元素
	f(intSlice)
	// 切片的类型必须匹配（下面编译不通过）
	// f(intSlice...)
	// 不传参 和 传 nil 的区别
	f()    // 0
	f(nil) // 1

}

func defaultVarargsBestPractice() {
	good1 := func(def string, args ...string) string {
		return append(args, def)[0]
	}
	good2 := func(args ...string) string {
		def := "123"
		return append(args, def)[0]
	}

	_, _ = good1, good2
}
