package benchmark

import (
	"encoding/json"
	"fmt"
	"testing"
)

// go test -bench=ToString$ -run=none .
// goos: windows
// goarch: amd64
// pkg: go-learn/special/performance/benchmark
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// Benchmark_FmtToString-8          6575601               204.8 ns/op
// Benchmark_JsonToString-8         4722921               271.8 ns/op
// PASS
// ok      go-learn/special/performance/benchmark  3.125s

// json 还没有将 []byte → string 就已经输了

func Benchmark_FmtToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(1.02)
	}
}

func Benchmark_JsonToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(1.02)
	}
}