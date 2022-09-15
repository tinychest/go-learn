package _05construct_binary_tree_from_preorder_and_inorder_traversal

import (
	. "go-learn/special/algorithm/leetcode/0common/tree_node"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case01",
			args: args{
				preorder: []int{1, 2, 4, 5, 3, 6, 7},
				inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			},
			want: FromArr(Ptr(1), Ptr(2), Ptr(3), Ptr(4), Ptr(5), Ptr(6), Ptr(7)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
