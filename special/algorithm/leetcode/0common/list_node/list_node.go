package list_node

import (
	"strconv"
	"strings"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func FromArr(nums ...int) *ListNode {
	var head = new(ListNode)
	p := head
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return head.Next
}

func (l *ListNode) String() string {
	p := l

	arr := make([]string, 0)
	for p != nil {
		arr = append(arr, strconv.Itoa(p.Val))
		p = p.Next
	}
	return "[" + strings.Join(arr, " → ") + "]"
}
