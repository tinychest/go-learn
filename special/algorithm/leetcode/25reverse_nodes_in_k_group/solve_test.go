package _5reverse_nodes_in_k_group

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				k:    2,
			},
			want: FromArr(2, 1, 4, 3, 5),
		},
		{
			name: "case02",
			args: args{
				head: FromArr(1, 2, 3, 4, 5),
				k:    3,
			},
			want: FromArr(3, 2, 1, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
