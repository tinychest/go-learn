package benchmark

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

/* 参照：https://geektutu.com/post/hpg-string-concat.html */

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 直接拼接
func directConcat(str string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// fmt.Sprintf
func sprintfConcat(str string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", str, s)
	}
	return s
}

// strings.Builder
func builderConcat(str string, n int) string {
	b := strings.Builder{}
	for i := 0; i < n; i++ {
		b.WriteString(str)
	}
	return b.String()
}

func preBuilderConcat(str string, n int) string {
	b := strings.Builder{}
	b.Grow(len(str) * n)
	for i := 0; i < n; i++ {
		b.WriteString(str)
	}
	return b.String()
}

// bytes.Buffer
func bufferConcat(str string, n int) string {
	b := bytes.Buffer{}
	for i := 0; i < n; i++ {
		b.WriteString(str)
	}
	return b.String()
}

func preBufferConcat(str string, n int) string {
	b := bytes.Buffer{}
	b.Grow(len(str) * n)
	for i := 0; i < n; i++ {
		b.WriteString(str)
	}
	return b.String()
}

// 字节切片
func bytesConcat(str string, n int) string {
	bs := make([]byte, 0)
	for i := 0; i < n; i++ {
		bs = append(bs, str...)
	}
	return string(bs)
}

// 空间预分配的字节切片
func preBytesConcat(str string, n int) string {
	bs := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		bs = append(bs, str...)
	}
	return string(bs)
}

func benchmark(b *testing.B, f func(string, int) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(str, 10000)
	}
}

// 性能测试命令：
// go test -bench="Concat$" -run=none -benchmem .
// 结果如下：
// Benchmark_DirectConcat-8               13          84428500 ns/op        530996238 B/op     10009 allocs/op
// Benchmark_SprintfConcat-8               7         150704214 ns/op        833453090 B/op     40873 allocs/op
// Benchmark_BuilderConcat-8            6315            162987 ns/op          505840 B/op         24 allocs/op
// Benchmark_PreBuilderConcat-8        16269             73893 ns/op          106496 B/op          1 allocs/op
// Benchmark_BufferConcat-8             7857            161259 ns/op          423536 B/op         13 allocs/op
// Benchmark_PreBufferConcat-8         10000            107132 ns/op          212992 B/op          2 allocs/op
// Benchmark_BytesConcat-8             10000            155849 ns/op          612337 B/op         25 allocs/op
// Benchmark_PreBytesConcat-8          15356             78505 ns/op          212992 B/op          2 allocs/op
// 结论：
// - 直接拼串 和 fmt.Sprintf 时间 和 空间 消耗都是最猛的
// - 推荐使用预分配内容的 strings.Builder 这是性能最好的（比预分配内存的 []byte 少了一次内存空间的转化）
func Benchmark_DirectConcat(b *testing.B)     { benchmark(b, directConcat) }
func Benchmark_SprintfConcat(b *testing.B)    { benchmark(b, sprintfConcat) }
func Benchmark_BuilderConcat(b *testing.B)    { benchmark(b, builderConcat) }
func Benchmark_PreBuilderConcat(b *testing.B) { benchmark(b, preBuilderConcat) }
func Benchmark_BufferConcat(b *testing.B)     { benchmark(b, bufferConcat) }
func Benchmark_PreBufferConcat(b *testing.B)  { benchmark(b, preBufferConcat) }
func Benchmark_BytesConcat(b *testing.B)      { benchmark(b, bytesConcat) }
func Benchmark_PreBytesConcat(b *testing.B)   { benchmark(b, preBytesConcat) }
