package benchmark

import "testing"

/* 参见：https://gfw.go101.org/article/value-copy-cost.html */

type S [12]int64
var sX = make([]S, 1000)
var sY = make([]S, 1000)
var sZ = make([]S, 1000)
var sumX, sumY, sumZ int64

func Benchmark_For(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumX = 0
		for j := 0; j < len(sX); j++ {
			sumX += sX[j][0]
		}
	}
}

func Benchmark_For_Range_OneIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumY = 0
		for j := range sY {
			sumY += sY[j][0]
		}
	}
}

func Benchmark_For_Range_TwoIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumZ = 0
		for _, v := range sZ {
			sumZ += v[0]
		}
	}
}

// [命令]
// go test -bench="Concat$" -benchmem .

// [结果]
// goos: windows
// goarch: amd64
// pkg: go-learn/special/performance/benchmark
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// Benchmark_For-8                           727695              1731 ns/op               0 B/op          0 allocs/op
// Benchmark_For_Range_OneIterVar-8          684794              1768 ns/op               0 B/op          0 allocs/op
// Benchmark_For_Range_TwoIterVar-8          307290              3723 ns/op               0 B/op          0 allocs/op
// PASS
// ok      go-learn/special/performance/benchmark  5.687s

// [结论]
// 数组这样的比较大的值类型（不仅是数组，还有字段比较多的结构体），拷贝的成本是很高的，实际操作尽量通过下标去操作