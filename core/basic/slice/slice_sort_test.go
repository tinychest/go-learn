package slice

import (
	"sort"
	"testing"
)

// sort 包默认按照 从小到大 的顺序进行排序

func TestSort(t *testing.T) {
	intSlice := []int{1, 2, 3}
	float64Slice := []float64{1, 2, 3}
	stringSlice := []string{"1", "2", "3"}

	// [基础类型的排序]
	sort.Ints(intSlice)
	sort.Float64s(float64Slice)
	sort.Strings(stringSlice)

	// [泛用排序方法]
	// sort.Sort(Interface)
	// sort.Slice(x any, less func(i, j int) bool)
	// - 要求稳定的泛用排序方法
	// sort.Stable(Interface)
	// sort.SliceStable(x any, less func(i, j int) bool)
	sort.Sort(sort.IntSlice(intSlice))
	sort.Sort(sort.Float64Slice(float64Slice))
	sort.Sort(sort.StringSlice(stringSlice))

	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] < intSlice[j]
	})
	sort.Slice(float64Slice, func(i, j int) bool {
		return float64Slice[i] < float64Slice[j]
	})
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i] < stringSlice[j]
	})

	// [从大到小的逆序]
	// Interface sort.Reverse(Interface)
}
