package _06reverse_linked_list

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
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
			},
			want: FromArr(5, 4, 3, 2, 1),
		},
		{
			name: "case02",
			args: args{
				head: FromArr(1, 2),
			},
			want: FromArr(2, 1),
		},
		{
			name: "case03",
			args: args{
				head: FromArr(),
			},
			want: FromArr(),
		},
		{
			name: "case04",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
