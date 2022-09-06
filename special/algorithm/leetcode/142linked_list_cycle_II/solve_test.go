package _42linked_list_cycle_II

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

func buildCycleSpecial(arr []int, cycleIndex int, cycleNode *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	var node *ListNode
	for i, v := range arr {
		if i == cycleIndex {
			node = cycleNode
		} else {
			node = &ListNode{Val: v, Next: nil}
		}

		p.Next = node
		p = node
	}

	p.Next = cycleNode
	return head.Next
}

func Test_detectCycle(t *testing.T) {
	special := &ListNode{Val: 2}
	args := buildCycleSpecial([]int{1, 2, 3, 4}, 1, special)
	want := special
	if got := detectCycle(args); !reflect.DeepEqual(got, want) {
		t.Errorf("case01 detectCycle() = %v, want %v", got, want)
	}

	special = nil
	args = FromArr(1)
	want = special
	if got := detectCycle(args); !reflect.DeepEqual(got, want) {
		t.Errorf("case02 detectCycle() = %v, want %v", got, want)
	}
}
