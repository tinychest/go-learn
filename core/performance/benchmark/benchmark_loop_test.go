package benchmark

import (
	"strings"
	"testing"
)

/* 参见：https://gfw.go101.org/article/value-copy-cost.html */

// [命令]
// go test -bench="Concat$" -benchmem .
//
// [结果]
// goos: windows
// goarch: amd64
// pkg: go-learn/special/performance/benchmark
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// Benchmark_For_Idx-8                       727695              1731 ns/op               0 B/op          0 allocs/op
// Benchmark_For_Range_OneIterVar-8          684794              1768 ns/op               0 B/op          0 allocs/op
// Benchmark_For_Range_TwoIterVar-8          307290              3723 ns/op               0 B/op          0 allocs/op
// PASS
// ok      go-learn/special/performance/benchmark  5.687s
//
// [结论]
// 数组这样的比较大的值类型（不仅是数组，还有字段比较多的结构体），拷贝的成本是很高的，实际操作尽量通过下标去操作

type arr [12]int64

var i1 = make([]arr, 1000)
var i2 = make([]arr, 1000)
var i3 = make([]arr, 1000)
var sum1, sum2, sum3 int64

func Benchmark_For_Idx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum1 = 0
		for j := 0; j < len(i1); j++ {
			sum1 += i1[j][0]
		}
	}
}

func Benchmark_For_Range_OneIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum2 = 0
		for j := range i2 {
			sum2 += i2[j][0]
		}
	}
}

func Benchmark_For_Range_TwoIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum3 = 0
		for _, v := range i3 {
			sum3 += v[0]
		}
	}
}

/* 对不同方式的 string 遍历进行性能基准测试 */

// [命令]
// go test -bench="For_String" -benchmem -run="none" -v

const size = 100

var s = randomString(size)
var bs = []byte(s)
var sb = strings.Builder{}

func Benchmark_For_String_Idx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sb.Reset()
		sb.Grow(size)
		for j := 0; j < size; j++ {
			if err := sb.WriteByte(bs[j]); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func Benchmark_For_String_Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sb.Reset()
		sb.Grow(size)
		for _, v := range bs {
			if err := sb.WriteByte(v); err != nil {
				b.Fatal(err)
			}
		}
	}
}
