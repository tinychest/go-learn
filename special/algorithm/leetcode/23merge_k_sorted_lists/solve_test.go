package _3merge_k_sorted_lists

import (
	"encoding/json"
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case01",
			args: args{
				lists: []*ListNode{
					FromArr(1, 4, 5),
					FromArr(1, 3, 4),
					FromArr(2, 6),
				},
			},
			want: FromArr(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			name: "case02",
			args: args{
				lists: []*ListNode{},
			},
			want: nil,
		},
		{
			name: "case03",
			args: args{
				lists: []*ListNode{nil},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
