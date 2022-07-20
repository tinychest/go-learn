package benchmark

import (
	"reflect"
	"testing"
)

// go test -bench="Nil" -benchmem .
//
// goos: windows
// goarch: amd64
// pkg: go-learn/unit_test/benchmark
// BenchmarkNil-8          177514371                6.80 ns/op            0 B/op          0 allocs/op
// BenchmarkNilStr-8       1000000000               0.310 ns/op           0 B/op          0 allocs/op
// PASS
// ok      go-learn/unit_test/benchmark    2.425s
//
// 结论：这恐怖的性能差距
func Benchmark_Nil(b *testing.B) {
	var strPtr *string

	for i := 0; i < b.N; i++ {
		IsNil(strPtr)
	}
}

func Benchmark_NilStr(b *testing.B) {
	var strPtr *string

	for i := 0; i < b.N; i++ {
		IsStrPtrNil(strPtr)
	}
}

func IsNil(param interface{}) bool {
	return reflect.ValueOf(param).IsNil()
}

func IsStrPtrNil(strPtr *string) bool {
	return strPtr == nil
}
