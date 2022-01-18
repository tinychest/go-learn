package http_rpc

import (
	"go-learn/core/net/rpc"
	"testing"
)

// 对比 http 请求接口调用、rpc 接口调用（每次重新建立连接并断开连接）、rpc 接口调用（连接复用）

// ❯ go test -bench='Call$' -run=none .
// goos: windows
// goarch: amd64
// pkg: go-learn/special/performance/benchmark/http_rpc
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// Benchmark_HTTPCall-8                9144            128996 ns/op
// Benchmark_RPCCall-8                 1378            819248 ns/op
// Benchmark_RPCReuseCall-8           14466             79634 ns/op
// PASS
// ok      go-learn/special/performance/benchmark/http_rpc 5.658s

func Benchmark_HTTPCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.HTTPClient()
	}
}

func Benchmark_RPCCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.RPCClient()
	}
}

func Benchmark_RPCReuseCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.RPCClientReuse()
	}
}