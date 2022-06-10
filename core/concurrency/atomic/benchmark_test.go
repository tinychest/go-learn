package atomic

import (
	"sync"
	"testing"
)

// 跑一个基准测试，看下两种同步方式上得性能差异
//
// go test -bench="^Benchmark" -run=none -benchmem .
//
// goos: windows
// goarch: amd64
// pkg: go-learn/core/concurrency/atomic
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// BenchmarkNormal-8            416           2446831 ns/op          240245 B/op     10002 allocs/op
// BenchmarkAtomic-8            476           3102951 ns/op          240028 B/op     10002 allocs/op
// BenchmarkMutex-8             367           3326102 ns/op          320142 B/op     10004 allocs/op
// PASS
// ok      go-learn/core/concurrency/atomic        4.647s
//
// Atomic 和 Mutex 之间对比，始终是 Atomic 占优（效果不明显）

func TestMutex(t *testing.T) {
	t.Log(MutexTest()) // = 10000
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalTest()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AtomicTest()
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MutexTest()
	}
}

func MutexTest() int64 {
	wg := sync.WaitGroup{}
	wg.Add(N)

	m := sync.Mutex{}

	var sum int64
	for i := 0; i < N; i++ {
		go func() {
			m.Lock()
			sum++
			m.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()
	return sum
}
