package _1merge_two_sorted_lists

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

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
				l1: FromArr(1, 2, 4),
				l2: FromArr(1, 3, 4),
			},
			want: FromArr(1, 1, 2, 3, 4, 4),
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
				l2: FromArr(1),
			},
			want: FromArr(1),
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
