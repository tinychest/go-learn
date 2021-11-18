package slice

import (
	"go-learn/util"
	"testing"
)

// 切片示例的切片声明不能超过父切片的上限，但是增加元素自动增长可以的
func TestSliceToSlice(t *testing.T) {
	intSlice := []int{1, 2, 3}

	tempSlice := intSlice[:1:2]
	util.PrintSlice(tempSlice)

	// 这个 append 直接影响了父切片里的元素
	tempSlice = append(tempSlice, 4)

	util.PrintSlice(tempSlice)
	util.PrintSlice(intSlice)
}

// 其实还是 append 和 切片的原理规则，当只有当 append 元素后，元素的数量大于 capacity，append 才会重新申请一个内存空间
// 才会脱离原来的切片，没有脱离父切片的子切片，其中一方操作对应位置的元素，都会影响另外一方
func test1() {
	intSlice := []int{1, 2, 3}

	tempSlice := intSlice[:0:0]
	// util.PrintSlice(tempSlice)
	// util.PrintSlice([]int{})
	// // 很特别的地址：0x0 length:0 capacity:0
	// util.PrintSlice(*new([]int))

	// 添加元素后切片元素个数为 1，大于切片的 capacity，所有会重新申请内存空间，创建一个新的切片
	tempSlice = append(tempSlice, 4)
	tempSlice = append(tempSlice, 5)
	tempSlice = append(tempSlice, 6)
	util.PrintSlice(intSlice)
	util.PrintSlice(tempSlice)
}
