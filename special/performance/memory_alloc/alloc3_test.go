package main

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"strconv"
	"testing"
)

// 优化3：很明显，下面的循环有一个 strings.Repeat 可以优化，分析可以发现，你的 x 想在 26 个字母中取一个，不应该重复这 26 个字母，而应该对这个下标取模
// 159 → 23
func foo3(n int) string {
	var buf bytes.Buffer

	x := strconv.Itoa(n)
	for i := 0; i < 100000; i++ {
		buf.WriteString(x)
	}
	sum := sha256.Sum256(buf.Bytes())

	var b []byte
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := "abcdefghijklmnopqrstuvwxyz"[x%26]
		b = append(b, c)
	}
	return string(b)
}

func TestAlloc3(t *testing.T) {
	t.Log("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo3(rand.Int())
	})))
}
