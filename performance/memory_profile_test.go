package performance

import (
	"github.com/pkg/profile"
	"math/rand"
	"testing"
)

// 这里内存性能分析，并没有使用 Go 的原生类库 runtime/pprof，而是一个易用性更强的库 pkg/profile，它封装了 runtime/pprof 的接口，使用起来更简单

// 辅助方法
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 实际测试的方法
func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func BenchmarkMemoryProfiling(b *testing.B) {
	// defer profile.Start().Stop() // 只要这一行代码，就能拿到接下来代码的 cpu 性能分析文件
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop() // 拿到 memory 的性能分析文件
	concat(100)
}
