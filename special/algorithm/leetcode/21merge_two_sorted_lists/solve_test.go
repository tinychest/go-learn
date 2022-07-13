package _1merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func nodesFromArr(arr ...int) *ListNode {
	if len(arr) == 0 {
		panic("oh dear!")
	}
	var head = &ListNode{Val: arr[0]}

	p := head
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i]}
		p.Next = node
		p = node
	}
	return head
}

func Test_mergeTwoLists(t *testing.T) {
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
				l1: nodesFromArr(1, 2, 4),
				l2: nodesFromArr(1, 3, 4),
			},
			want: nodesFromArr(1, 1, 2, 3, 4, 4),
		},
		{
			name: "case02",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			name: "case03",
			args: args{
				l1: nil,
				l2: nodesFromArr(1),
			},
			want: nodesFromArr(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
