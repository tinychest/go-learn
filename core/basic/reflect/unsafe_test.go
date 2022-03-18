package _reflect

import (
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	// conceptTest(t)
	// sliceElemOffsetTest(t)
	// sliceChildTest(t)
}

func conceptTest(t *testing.T) {
	// unsafe.Pointer → *ArbitraryType → *int
	// _ = uintptr(1) 语法通过

	// *任意类型 ←→ unsafe.Pointer
	p := unsafe.Pointer(nil) // 可以包裹任意 * 类型
	_ = (*int)(p)            // unsafe.Pointer 可以被断言成任何类型，实际不是就会 panic；即使显而易见不对，也不会基于提示，所以是 unsafe

	// unsafe.Pointer → uintptr →（会有警告） unsafe.Pointer
	addr := uintptr(p)       // unsafe.Pointer 可以转成地址值
	p = unsafe.Pointer(addr) // 将任意一个地址值，包装成 unsafe.Pointer 去使用，编译器会给出警告
}

func sliceElemOffsetTest(t *testing.T) {
	s := [4]int{0: 0, 1: 1, 2: 2, 3: 3}
	elemSize := int(unsafe.Sizeof(s[0]))
	e4 := &s[3]
	ep4 := unsafe.Pointer(e4)
	e3 := (*int)(unsafe.Add(ep4, -1*elemSize))
	t.Log(*e3) // 2

	e5 := unsafe.Add(ep4, elemSize) // 引用了未知的内存块，实际运行都不一定报错，这就是 unsafe
	_ = e5
}

func sliceChildTest(t *testing.T) {
	s := [4]int{0: 0, 1: 1, 2: 2, 3: 3}
	e2 := &s[1]

	// 从 s 为起始地址，向后取 3 个元素（cap），实际 2（len）
	s1 := unsafe.Slice(e2, 3)
	t.Log(s1)               // [1 2 3]
	t.Log(len(s1), cap(s1)) // 3 3
	s2 := s1[:2]
	t.Log(s2)               // [1 2]
	t.Log(len(s2), cap(s2)) // 2 3

	_ = unsafe.Slice(e2, 4) // 引用了未知的内存块，实际运行都不一定报错，这就是 unsafe
}
