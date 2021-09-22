package benchmark

import (
	"math/rand"
	"testing"
	"time"
)

/*
测试一：当前文件的 benchmark
go test -bench="IntSlice$" .
goos: windows
goarch: amd64
pkg: go-learn/unit_test/benchmark
BenchmarkForIntSlice-8              3432            332459 ns/op
BenchmarkRangeIntSlice-8            3750            334121 ns/op
PASS
ok      go-learn/unit_test/benchmark    2.802s

结论：多次测试各有千秋，可以认为性能没有差别

测试二：因为 for range 的遍历过程实际上会对 value 进行一次值拷贝，所以当 for range 的是值，且特别大时，for 下标遍历值和 for range 遍历值将会有较大的性能差距
（只遍历下标，没有特别大的区别）

测试三：测试二中的大值类型，换为指针类型（[]struct → []*struct），遍历值将不会有什么区别

总结：range 在迭代过程中返回的是迭代值的拷贝，如果每次迭代的元素的内存占用很低，那么 for 和 range 的性能几乎是一样，例如 []int。
但是如果迭代的元素内存占用较高，例如一个包含很多属性的 struct 结构体，那么 for 的性能将显著地高于 range，有时候甚至会有上千倍的性能差异。
对于这种场景，建议使用 for，如果使用 range，建议只迭代下标，通过下标访问迭代值，这种使用方式和 for 就没有区别了。
如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，才能不影响性能。
*/

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkForIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)

	for i := 0; i < b.N; i++ {
		len := len(nums)
		var tmp int
		for k := 0; k < len; k++ {
			tmp = nums[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)

	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range nums {
			tmp = num
		}
		_ = tmp
	}
}
