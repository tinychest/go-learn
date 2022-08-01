package slice

import (
	"testing"
)

// - 对于 for range 中对切片元素进行字段修改，不是一下蒙头进哪个哪个结构体里的字段是不是指针类型，会不会有不同的影响，而是要从一个大的全局出发
// - 在 for range 中对遍历的切片进行 append，并不会引起死循环；可以理解为在执行 for range 的时候就确定了切片，不会因后续的变动收到影响
func TestForRangeEffectToSlice(t *testing.T) {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		t.Log(v)
		arr = append(arr, 4)
	}
}
