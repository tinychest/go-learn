package _9remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"
)

func nodesFromArr(arr []int) *ListNode {
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

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case01",
			args: args{
				head: nodesFromArr([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: nodesFromArr([]int{1, 2, 3, 5}),
		},
		{
			name: "case02",
			args: args{
				head: nodesFromArr([]int{1}),
				n:    1,
			},
			want: nil,
		},
		{
			name: "case03",
			args: args{
				head: nodesFromArr([]int{1, 2}),
				n:    1,
			},
			want: nodesFromArr([]int{1}),
		},
		// submit error → 定义一个头指针来解决这个
		{
			name: "error-case01",
			args: args{
				head: nodesFromArr([]int{1, 2}),
				n:    2,
			},
			want: nodesFromArr([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
