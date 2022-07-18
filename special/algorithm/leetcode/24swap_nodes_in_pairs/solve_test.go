package _4swap_nodes_in_pairs

import (
	"encoding/json"
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
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
				head: FromArr(1, 2, 3, 4),
			},
			want: FromArr(2, 1, 4, 3),
		},
		{
			name: "case02",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "case03",
			args: args{
				head: FromArr(1),
			},
			want: FromArr(1),
		},
		{
			name: "case04",
			args: args{
				head: FromArr(1, 2, 3, 4, 5, 6),
			},
			want: FromArr(2, 1, 4, 3, 6, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
