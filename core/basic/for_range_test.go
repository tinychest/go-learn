package basic

import (
	"fmt"
	"testing"
)

// 对于 map 来说
// 1.和切片不同的是，迭代过程中，删除还未迭代到的键值对，则该键值对不会被迭代
// 2.在迭代过程中，如果创建新的键值对，那么新增键值对，可能被迭代，也可能不会被迭代。

// 对于 可读通道 来说
// 1.只要通道没有关闭，for range 就不会结束
// 2.如果通道为 nil，for range 将被永远阻塞
func TestForRange(t *testing.T) {
	intSlice := []int{1, 2, 3}

	// 0xc00000e480
	// 0xc00000e488
	// 0xc00000e490
	for index := 0; index < len(intSlice); index++ {
		fmt.Printf("%p\n", &intSlice[index])
	}

	// 0xc00000a320
	// 0xc00000a320
	// 0xc00000a320
	for _, intValue := range intSlice {
		fmt.Printf("%p\n", &intValue)
	}
}
