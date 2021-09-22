package util

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.In(1).In(2).In(3)

	for !q.IsEmpty() {
		elem, _ := q.Out()
		fmt.Print(strconv.Itoa(elem.(int)) + " ")
	}
	fmt.Println()

	// 2
	q = NewQueue()
	for !q.IsEmpty() {
		elem, _ := q.Out()
		fmt.Print(strconv.Itoa(elem.(int)) + " ")
	}
	fmt.Println()

	// 3
	fmt.Println(NewQueue().Out())
}
