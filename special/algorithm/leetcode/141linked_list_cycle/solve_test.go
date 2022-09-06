package _41linked_list_cycle

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"testing"
)

func buildCycle(arr []int, cycleIndex int) *ListNode {
	head := &ListNode{}
	p := head

	var node *ListNode
	special := &ListNode{}
	for i, v := range arr {
		if i == cycleIndex {
			special.Val = v
			node = special
		} else {
			node = &ListNode{Val: v, Next: nil}
		}

		p.Next = node
		p = node
	}

	p.Next = special
	return head.Next
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case01",
			args: args{
				head: FromArr(3, 2, 0, -1),
			},
			want: false,
		},
		{
			name: "case02",
			args: args{
				head: buildCycle([]int{3, 2, 0, -1}, 2),
			},
			want: true,
		},
		{
			name: "case03",
			args: args{
				head: buildCycle([]int{0}, 0),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
