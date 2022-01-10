package benchmark

import (
	"encoding/json"
	"go-learn/core"
	"sync"
	"testing"
)

// sync.Pool 是什么，源码注释上写的非常清楚了，其中比较核心的一句
// Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector

// go test -bench='New$' -benchmem .
// goos: windows
// goarch: amd64
// pkg: go-learn/special/performance/benchmark
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// Benchmark_New-8                   2938779               398.4 ns/op           176 B/op          3 allocs/op
// Benchmark_PoolNew-8               2155370               607.3 ns/op           194 B/op          3 allocs/op

// 似乎并没有发现特别大的差异，无论是在时间上，还是空间上

var bs = []byte("{}")

var personPool = sync.Pool{
	New: func() interface{} {
		return new(core.Person)
	},
}

func Benchmark_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := new(core.Person)
		_ = json.Unmarshal(bs, p)
	}
}

func Benchmark_PoolNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := personPool.Get().(*core.Person)
		_ = json.Unmarshal(bs, p)
		studentPool.Put(p)
	}
}
