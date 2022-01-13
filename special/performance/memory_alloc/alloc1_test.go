package main

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

// 优化1：为 buf 预设缓存大小 100158 → 100148（就只减少了 10 次...）
func foo1(n int) string {
	var buf bytes.Buffer
	for i := 0; i < 100000; i++ {
		buf.WriteString(strconv.Itoa(n))
	}
	sum := sha256.Sum256(buf.Bytes())

	var b []byte
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10)[x]
		b = append(b, c)
	}
	return string(b)
}

func TestAlloc1(t *testing.T) {
	t.Log("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo1(rand.Int())
	})))
}
