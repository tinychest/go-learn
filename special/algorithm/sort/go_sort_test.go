package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	intSlice := sort.IntSlice{3, 2, 1}

	// 方式1
	intSlice.Sort()
	fmt.Println(intSlice)

	// 方式2
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] <= intSlice[j]
	})
	fmt.Println(intSlice)

	// 方式3（自定义类型，只要实现 sort.Interface 即可，方式1 本质就是这个）
	// sort.Sort()
}