package tool

import (
	"container/list"
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
}

// 标准库 list
func TestList(t *testing.T) {
	l := list.New()
	// l.PushBack()
	// l.PushFront()
	// l.Remove()
	// l.Len()
	// l.Init() // clear
	// l.Back() // latest elem

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// t.Log(l.Back().Value)
	// t.Log(l.Back().Prev().Value)

	for v := l.Front(); v != nil; v = v.Next() {
		t.Log(v.Value)
	}
}
