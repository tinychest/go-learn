package _iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	intSlice := IntSlice{1, 2, 3}

	iterator := intSlice.Iterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
