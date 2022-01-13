package sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{1, 2, 45, 12, 11, 21, 0}

	InsertSort(arr)

	t.Log(arr)
}
