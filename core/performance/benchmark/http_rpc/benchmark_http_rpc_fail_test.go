package http_rpc

// go run http/http_server.go
// go run http/rpc_server.go
// go test -bench='Call$' -run=none .

// benchmark_http_rpc_test.go:4:2: import "go-learn/special/performance/benchmark/http_rpc/http" is a program, not an importable package
// FAIL    go-learn/special/performance/benchmark/http_rpc [setup failed]
// FAIL
// （GG 虽然 Goland 没有提示错误，实际凉凉）

// import (
// 	http "go-learn/special/performance/benchmark/http"
// 	rpc "go-learn/special/performance/benchmark/rpc"
// 	"testing"
// )

// func Benchmark_HTTPCall(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		http.HTTPClient()
// 	}
// }
//
// func Benchmark_RPCCall(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		rpc.RPCClient()
// 	}
// }
