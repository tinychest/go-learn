package search

import (
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	// 搜索 0 和 1 返回的都是 0
	index0 := sort.SearchInts([]int{1, 3, 5, 7, 9}, 0)
	index1 := sort.SearchInts([]int{1, 3, 5, 7, 9}, 1)
	t.Log(index0)
	t.Log(index1)
}
