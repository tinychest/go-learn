package slice

import (
	"go-learn/tool"
	"testing"
)

// 切片性能陷阱（内存泄漏）
// 切片经由一通操作后，可能扩容的很大，之后假如要用到该切片的 几个元素，就不要继续使用该切片了（复制元素到新的小切片），占用太大空间得不到释放
// 参见 https://mp.weixin.qq.com/s/dejNOtGwID6z9ExLLybDtA

func TestBasic(t *testing.T) {
	// addressConcept(t)
	// expandConcept(t)

	childTest(t)
}

// 使用概念 没有初始化也是可以直接 append 使用的
func useConcept(t *testing.T) {
	var s []string

	tool.PrintSlice(s)
	s = append(s, "abc")
	tool.PrintSlice(s)
}

// 扩容概念
func expandConcept(t *testing.T) {
	// runtime/slice.go - growslice
	// - 旧切片长度小于 1024，最终容量会是旧容量的两倍
	// - 旧切片长度大于等于 1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的 1/4，直到大于或等于新申请的容量
	// - 最终得出的容量（cap）值已处于最大值，则限定为最大值

	// 0 → 1 → 2 → 4 → 8 → 16...

	var s []int
	// 注意添加元素，一定要通过循环去加，这样可测不出来，因为 Go 编译器有预处理
	// s = append(s, 1, 2, 3)
	for i := 0; i < 16; i++ {
		s = append(s, 1)
		tool.PrintSlice(s)
	}
}

// 内存地址的概念测试
// 测试0：切片 [零值] 和 [初始化] 的地址
// - 切片的本质是指向底层数组的指针，但是切片类型的变量的零值不是 nil
// - 切片的零值不是 nil，但是切片类型的变量可以赋值为 nil
func addressConcept(t *testing.T) {
	var s1 []int
	t.Logf("%p\n", s1)             // 0x0
	t.Logf("%p\n", []int(nil))     // 0x0
	t.Logf("%p\n", *new([]int))    // 0x0
	t.Logf("%p\n", new([]int))     // 地址1
	t.Logf("%p\n", []int{})        // 地址2
	t.Logf("%p\n", make([]int, 0)) // 地址3

	s2 := []int{1, 2, 3, 4, 5}
	cs2 := s2

	// 测试1：两个指针类型指向的数组的地址，在内存中是相同的（不会进行底层数组的拷贝）
	t.Logf("%p\n", s2)
	t.Logf("%p\n", cs2)

	// 测试2：从第2个元素开始的切片的地址和底层数组第2个元素的地址是相同的（同 C 中数组的地址等同于数组第一个元素的地址）
	t.Logf("%p\n", &s2[1])
	t.Logf("%p\n", cs2[1:])

	// 测试3：方法传参传的也是实际地址值
	// 将切片作为函数参数，并在函数修改切片的值，是能够真实影响数组的值（传给函数，并不会将底层数组复制一个副本），但是：数组是不行的！！！
	s3 := []int{1, 2, 3}
	t.Logf("%p\n", s3)

	func(s []int) {
		t.Logf("%p\n", s)
	}(s3)
}

func childTest(t *testing.T) {
	slice := []int{1, 2, 3}

	// 取头舍尾，所以实际是 slice[1:2] = [2]
	t.Log(slice[1 : len(slice)-1])

	// 子切片的定义不能超出父切片的范围
	// t.Log(slice[:4])

	t.Log(slice[2:])
}
