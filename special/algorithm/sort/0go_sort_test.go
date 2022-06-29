package sort

import (
	"sort"
	"testing"
)

// TODO 简述一下 Go 实现的细节

func TestSort(t *testing.T) {
	intSlice := sort.IntSlice{3, 2, 1}

	// 方式1
	intSlice.Sort()
	t.Log(intSlice)

	// 方式2
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] <= intSlice[j]
	})
	t.Log(intSlice)

	// 方式3（自定义类型，只要实现 sort.Interface 即可，方式1 本质就是这个）
	// sort.Sort()
}
