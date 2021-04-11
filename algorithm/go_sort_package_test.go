package algorithm

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	// 没什么卵用，搜索 0 和 1 返回的都是 0
	// 写法很简单
	index := sort.SearchInts([]int{1, 3, 5, 7, 9}, 1)
	println(index)
}
