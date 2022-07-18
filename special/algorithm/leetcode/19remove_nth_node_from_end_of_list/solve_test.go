package _9remove_nth_node_from_end_of_list

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

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
				head: FromArr(1, 2, 3, 4, 5),
				n:    2,
			},
			want: FromArr(1, 2, 3, 5),
		},
		{
			name: "case02",
			args: args{
				head: FromArr(1),
				n:    1,
			},
			want: nil,
		},
		{
			name: "case03",
			args: args{
				head: FromArr(1, 2),
				n:    1,
			},
			want: FromArr(1),
		},
		// submit error → 定义一个头指针来解决这个
		{
			name: "error-case01",
			args: args{
				head: FromArr(1, 2),
				n:    2,
			},
			want: FromArr(2),
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
