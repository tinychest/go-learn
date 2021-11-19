package _5three_sum

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2,3,4,1,2,5,12}
	quickSort(arr)
	fmt.Println(arr)
}

func TestHalfSearch(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7}
	fmt.Println(halfSearch(arr, 6))
	fmt.Println(halfSearch(arr, 8))
}
