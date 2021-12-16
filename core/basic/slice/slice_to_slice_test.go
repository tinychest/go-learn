package slice

import (
	"go-learn/util"
	"testing"
)

// 切片示例的切片声明不能超过父切片的上限，但是增加元素自动增长可以的
func TestQuota1(t *testing.T) {
	slice := []int{1, 2, 3}

	ref := slice[:1:2]
	util.PrintSlice(ref)

	// 这个 append 直接影响了父切片里的元素
	ref = append(ref, 4)
	util.PrintSlice(ref)
	util.PrintSlice(slice)
}

// 其实还是 append 和 切片的原理规则，当只有当 append 元素后，元素的数量大于 capacity，append 才会重新申请一个内存空间
// 才会脱离原来的切片，没有脱离父切片的子切片，其中一方操作对应位置的元素，都会影响另外一方
func TestQuota2(t *testing.T) {
	slice := []int{1, 2, 3}

	ref := slice[:0:0]
	// util.PrintSlice(tempSlice)
	// util.PrintSlice([]int{})
	// // 特殊的地址：0x0 length:0 capacity:0
	// util.PrintSlice(*new([]int))

	// 添加元素后切片元素个数为 1，大于切片的 capacity，所有会重新申请内存空间，创建一个新的切片
	ref = append(ref, 4)
	ref = append(ref, 5)
	ref = append(ref, 6)
	util.PrintSlice(slice)
	util.PrintSlice(ref)
}

// bytes.Buffer.Reset 方法中，重置底层字节数组是这样写的：b.buf = b.buf[:0]
func TestQuota3(t *testing.T) {
	s := make([]string, 0, 4)
	s = append(s, "1")
	util.PrintSlice(s)
	s = s[:0]
	util.PrintSlice(s)
	s = s[:1]
	util.PrintSlice(s) // 从这里的结果可以了解到，确实是复用空间，原来位置的值都没有变

	// 所以 bytes.Buffer.Write 的相关方法都是，从指定下标覆盖写
}
