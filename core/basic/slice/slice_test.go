package slice

import (
	"fmt"
	"testing"
)

// 切片性能陷阱（内存泄漏）
// 切片经由一通操作后，可能扩容的很大，之后假如要用到该切片的 几个元素，就不要继续使用该切片了（复制元素到新的小切片），占用太大空间得不到释放
// 参见 https://mp.weixin.qq.com/s/dejNOtGwID6z9ExLLybDtA

func TestBasic(t *testing.T) {
	// addressConcept()
	// expandConcept()

	childTest()
}

// 扩容概念
func expandConcept() {
	// runtime/slice.go - growslice
	// - 旧切片长度小于 1024，最终容量会是旧容量的两倍
	// - 旧切片长度大于等于 1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的 1/4，直到大于或等于新申请的容量
	// - 最终得出的容量（cap）值已处于最大值，则限定为最大值

	// 0 → 1 → 2 → 4 → 8 → 16...
}

// 内存地址的概念测试
func addressConcept() {
	// 测试0：切片 [零值] 和 [初始化] 的地址
	// - 切片的本质是指向底层数组的指针，但是切片类型的变量的零值不是 nil
	// - 切片的零值不是 nil，但是切片类型的变量可以赋值为 nil
	var intSlice []int
	fmt.Printf("%p\n", intSlice)    // 0x0
	fmt.Printf("%p\n", []int(nil))  // 0x0
	fmt.Printf("%p\n", *new([]int)) // 0x0

	fmt.Printf("%p\n", []int{})        // 地址1
	fmt.Printf("%p\n", make([]int, 0)) // 地址1

	fmt.Printf("%p\n", new([]int)) // 这个和上面的地址都不同，毕竟 数据类型 都不同嘛

	originSlice := []int{1, 2, 3, 4, 5}
	pointerSlice := originSlice

	// 测试1：两个指针类型指向的数组的地址，在内存中是相同的
	fmt.Printf("%p\n", originSlice)
	fmt.Printf("%p\n", pointerSlice)

	// 测试2：从第2个元素开始的切片的地址和底层数组第2个元素的地址是相同的（同 C 中数组的地址等同于数组第一个元素的地址）
	fmt.Printf("%p\n", &originSlice[1])
	fmt.Printf("%p\n", pointerSlice[1:])

	// 测试3：方法传参传的也是实际地址值
	// 将切片作为函数参数，并在函数修改切片的值，是能够真实影响数组的值（传给函数，并不会将底层数组复制一个副本），但是：数组是不行的！！！
	tempSlice := []int{1, 2, 3}
	fmt.Printf("%p\n", tempSlice)

	func(intSlice []int) {
		fmt.Printf("%p\n", intSlice)
	}(tempSlice)
}

func childTest() {
	slice := []int{1, 2, 3}

	// 取头舍尾，所以实际是 [2]
	// childSlice := slice[1:len(slice)-1]

	// 定义上，子切片的定义不能超出父切片的范围
	// println(childSlice[:2])

	fmt.Println(slice[2:])
}
