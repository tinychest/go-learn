package slice

import (
	"testing"
)

// copy 的依据只有 len
func TestCopy(t *testing.T) {
	source := []int{1, 2, 3}

	var target1 []int
	target2 := make([]int, 0)
	target3 := make([]int, 0, 3)
	target4 := make([]int, 3)

	t.Log(copy(target1, source)) // 0
	t.Log(copy(target2, source)) // 0
	t.Log(copy(target3, source)) // 0
	t.Log(copy(target4, source)) // 3
}
