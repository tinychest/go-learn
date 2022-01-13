package _5three_sum

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 3, 4, 1, 2, 5, 12}
	quickSort(arr)
	t.Log(arr)
}

func TestHalfSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(halfSearch(arr, 6))
	t.Log(halfSearch(arr, 8))
}
