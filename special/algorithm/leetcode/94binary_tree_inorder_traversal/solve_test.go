package _4binary_tree_inorder_traversal

import (
	. "go-learn/special/algorithm/leetcode/0common/tree_node"
	"reflect"
	"testing"
)

func ptr(i int) *int {
	return &i
}

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case01",
			args: args{
				root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "case02",
			args: args{
				root: FromArr(),
			},
			want: []int{},
		},
		{
			name: "case03",
			args: args{
				root: FromArr(ptr(1)),
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
