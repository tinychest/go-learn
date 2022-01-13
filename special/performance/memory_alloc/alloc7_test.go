package main

import (
	"crypto/sha256"
	"hash"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"unsafe"
)

// 优化7：使用 pool 重用 bytes.Buffer；重用 sha256
// sha256.New() hash.Hash 没有 WriteString 方法，所以需要借助 strconv.AppendInt 来避免额外的空间分配次数
// 3 → 1

// 重用 buf（首先将 bytes.Buffer 改成一个简单的 []byte）
var bufPool = sync.Pool{
	New: func() interface{} {
		// length of a sha256 hash
		b := make([]byte, 256)
		return &b
	},
}

// 重用
var hashPool = sync.Pool{
	New: func() interface{} { return sha256.New() },
}

func foo7(n int) string {
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

	var b = make([]byte, 0, int(sum[0]))
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := "abcdefghijklmnopqrstuvwxyz"[x%26]
		b = append(b, c)
	}

	return *(*string)(unsafe.Pointer(&b))
}

func TestAlloc7(t *testing.T) {
	t.Log("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo7(rand.Int())
	})))
}
