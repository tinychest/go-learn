package sort

import (
	"fmt"
	"testing"
)

func TestHalfSearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 7, 9, 22}

	index, ok := HalfSearch(11, array)

	fmt.Println(index, ok)
}
