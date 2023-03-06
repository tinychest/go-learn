package performance

import "testing"

// 在下面的 bench 性能测试场景，max 函数的调用会被内联优化，而加了禁止内联的标签的函数则不会
// ❯ go test -bench="nline" -run="none" .
// goos: windows
// goarch: amd64
// pkg: go-learn/core/performance
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// BenchmarkInline-8       1000000000               0.3231 ns/op
// BenchmarkNoinline-8     635348233                1.858 ns/op
// PASS
// ok      go-learn/core/performance       2.482s

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//go:noinline
func noinlineMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BenchmarkInline(b *testing.B) {
	x, y := 1, 2
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		max(x, y)
	}
}

func BenchmarkNoinline(b *testing.B) {
	x, y := 1, 2
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		noinlineMax(x, y)
	}
}
