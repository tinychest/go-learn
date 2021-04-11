package simple

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}

func BenchmarkSliceCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sliceCap(1000000)
	}
}

func BenchmarkSliceAppropriateCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sliceAppropriateCap(1000000)
	}
}
