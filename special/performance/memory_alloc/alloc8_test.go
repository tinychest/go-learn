package main

import (
	"fmt"
	"hash"
	"math/rand"
	"strconv"
	"testing"
	"unsafe"
)

// 优化8：最后一个内存分配在哪里？
// 很明显：var b = make([]byte, 0, int(sum[0]))
// 1 → 0

// 这里必须给出详细的解释，是如果避免最后 1 次内存分配的，以下是原文文章解释
// 如果我们将切片保留在函数中并且不返回它，它将被在栈上分配。
// 任何 C 程序员都知道，分配在栈上的内存是不可能返回的。
// 当 Go 编译器看到 return 语句时，切片会改为在堆上分配。
// 但这并不能真正解释为什么将内容传输到sum不分配。
// sum不是也在堆上分配了吗？是的，但sync.Pool已经为我们在堆上分配了
func foo8(n int) string {
	bufptr := bufPool.Get().(*[]byte)
	defer bufPool.Put(bufptr)
	buf := *bufptr
	buf = buf[:0]

	h := hashPool.Get().(hash.Hash)
	defer hashPool.Put(h)
	h.Reset()

	x := strconv.AppendInt(buf, int64(n), 10)
	for i := 0; i < 100000; i++ {
		h.Write(x)
	}
	sum := h.Sum(buf)

	var b = make([]byte, 0, 256)
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := "abcdefghijklmnopqrstuvwxyz"[x%26]
		b = append(b, c)
	}

	sum = sum[:0]
	sum = append(sum, b...)
	return *(*string)(unsafe.Pointer(&sum))
}

func TestAlloc8(t *testing.T) {
	fmt.Println("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo8(rand.Int())
	})))
}
