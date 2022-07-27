package _04maximum_depth_of_binary_tree

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{
				root: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: 3,
		},
		{
			name: "case02",
			args: args{
				root: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
			},
			want: 2,
		},
		{
			name: "case03",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: 1,
		},
		{
			name: "case03",
			args: args{
				root: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
