package basic

import (
	"testing"
)

// 对于 map 来说
// 1.和切片不同的是，迭代过程中，删除还未迭代到的键值对，则该键值对不会被迭代
// 2.在迭代过程中，如果创建新的键值对，那么新增键值对，可能被迭代，也可能不会被迭代。

// 对于 可读通道 来说
// 1.只要通道没有关闭，for range 就不会结束
// 2.如果通道为 nil，for range 将被永远阻塞
func TestForRange(t *testing.T) {
	s := []int{1, 2, 3}

	// 0xc00000e480
	// 0xc00000e488
	// 0xc00000e490
	for i := 0; i < len(s); i++ {
		t.Logf("%p\n", &s[i])
	}

	// 0xc00000a320
	// 0xc00000a320
	// 0xc00000a320
	for _, v := range s {
		// 注意，如果 v 是指针类型，且需要进行 append 操作，应该先 v := v
		t.Logf("%p\n", &v)
	}
}
