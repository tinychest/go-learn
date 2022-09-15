package _24binary_tree_maximum_path_sum

import (
	. "go-learn/special/algorithm/leetcode/0common/tree_node"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
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
				root: &TreeNode{Val: -2},
			},
			want: -2,
		},
		{
			name: "case02",
			args: args{
				root: FromArr(Ptr(1), Ptr(2), Ptr(3)),
			},
			want: 6,
		},
		{
			name: "case03",
			args: args{
				root: FromArr(Ptr(-10), Ptr(9), Ptr(20), nil, nil, Ptr(15), Ptr(7)),
			},
			want: 42,
		},
		{
			name: "case04",
			args: args{
				root: FromArr(Ptr(2), Ptr(-1), nil),
			},
			want: 2,
		},
		{
			name: "case05",
			args: args{
				root: FromArr(Ptr(-2), Ptr(-1), nil),
			},
			want: -1,
		},
		{
			name: "case06",
			args: args{
				root: FromArr(Ptr(1), Ptr(-2), Ptr(3)),
			},
			want: 4,
		},
		{
			name: "case07",
			args: args{
				root: FromArr(Ptr(1), Ptr(-2), Ptr(-3), Ptr(1), Ptr(3), Ptr(-2), nil, Ptr(-1)),
			},
			want: 3,
		},
		{
			name: "case08",
			args: args{
				root: FromArr(Ptr(5), Ptr(4), Ptr(8), Ptr(11), nil, Ptr(13), Ptr(4), Ptr(7), Ptr(2), nil, nil, nil, Ptr(1)),
			},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
