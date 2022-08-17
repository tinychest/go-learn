package _14flatten_binary_tree_to_linked_list

import (
	. "go-learn/special/algorithm/leetcode/0common/tree_node"
	"reflect"
	"testing"
)

func p(i int) *int { return &i }

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case01",
			args: args{
				root: FromArr(p(1), p(2), p(3)),
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "fail01",
			args: args{
				root: FromArr(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if got := tt.args.root; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
