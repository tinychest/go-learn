package _iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	intSlice := IntSlice{1, 2, 3}

	iterator := intSlice.Iterator()

	for iterator.HasNext() {
		t.Log(iterator.Next())
	}
}
