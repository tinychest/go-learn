package benchmark

import (
	"math/rand"
	"testing"
)

// 元素数量、元素的大小都会加大两种方式的差异
// 元素越多，元素越大 map 的性能优势越大

// go test -bench='Match$' -run=none .

const (
	elemSum = 10000
	elemLen = 10
)

func randomSlice(elemSum, elemLen int) []string {
	var s = make([]string, 0, elemSum)
	for i := 0; i < elemSum; i++ {
		s = append(s, randomString(elemLen))
	}
	return s
}

func BenchmarkMapMatch(b *testing.B) {
	s := randomSlice(elemSum, elemLen)

	var m = make(map[string]struct{}, elemSum)
	for i := 0; i < elemSum; i++ {
		m[s[i]] = struct{}{}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		elem := s[rand.Intn(len(s))]
		_ = m[elem]
	}
}

func BenchmarkSliceMatch(b *testing.B) {
	s := randomSlice(elemSum, elemLen)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		elem := s[rand.Intn(len(s))]
		for j := 0; j < len(s); j++ {
			if s[j] == elem {
				break
			}
		}
	}
}
