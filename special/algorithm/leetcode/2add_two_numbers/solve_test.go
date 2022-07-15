package _add_two_numbers

import (
	"encoding/json"
	"reflect"
	"testing"
)

func nodesFromArray(nums []int) *ListNode {
	var head = new(ListNode)
	p := head
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return head.Next
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case01",
			args: args{
				l1: nodesFromArray([]int{2, 4, 3}),
				l2: nodesFromArray([]int{5, 6, 4}),
			},
			want: nodesFromArray([]int{7, 0, 8}),
		},
		{
			name: "case02",
			args: args{
				l1: nodesFromArray([]int{0}),
				l2: nodesFromArray([]int{0}),
			},
			want: nodesFromArray([]int{0}),
		},
		{
			name: "case03",
			args: args{
				l1: nodesFromArray([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: nodesFromArray([]int{9, 9, 9, 9}),
			},
			want: nodesFromArray([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				bs1, _ := json.Marshal(got)
				bs2, _ := json.Marshal(tt.want)
				t.Log(string(bs1))
				t.Log(string(bs2))
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
