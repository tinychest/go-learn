package _01symmetric_tree

import "testing"

func ByVal(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case01",
			// 0
			args: args{ByVal(0)},
			want: true,
		},
		{
			name: "case02",
			// 0 1
			args: args{&TreeNode{Val: 0, Left: ByVal(1)}},
			want: false,
		},
		{
			name: "case03",
			// 0 1 1
			args: args{&TreeNode{
				Val:   0,
				Left:  ByVal(1),
				Right: ByVal(1),
			}},
			want: true,
		},
		{
			name: "case04",
			// 0 11 12
			args: args{&TreeNode{
				Val:   0,
				Left:  ByVal(11),
				Right: ByVal(12),
			}},
			want: false,
		},
		{
			name: "case05",
			// 0 1 1 2 nil nil 2
			args: args{&TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 1, Left: ByVal(2), Right: nil},
				Right: &TreeNode{Val: 1, Left: nil, Right: ByVal(2)},
			}},
			want: true,
		},
		{
			name: "case06",
			// 0 1 1 2 3 3 2
			args: args{&TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 1, Left: ByVal(2), Right: ByVal(3)},
				Right: &TreeNode{Val: 1, Left: ByVal(3), Right: ByVal(2)},
			}},
			want: true,
		},
		// [2,3,3,4,5,null,4]
		{
			name: "case07",
			// 0 1 1 2 3 nil 2
			args: args{&TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 1, Left: ByVal(2), Right: ByVal(3)},
				Right: &TreeNode{Val: 1, Left: nil, Right: ByVal(2)},
			}},
			want: false,
		},
		{
			name: "case08",
			// [0] [1 1] [2 2 2 2] [1 2 3 4 4 3 2 1]
			args: args{&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  ByVal(1),
						Right: ByVal(2),
					},
					Right: &TreeNode{
						Val:   2,
						Left:  ByVal(3),
						Right: ByVal(4),
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  ByVal(4),
						Right: ByVal(3),
					},
					Right: &TreeNode{
						Val:   2,
						Left:  ByVal(2),
						Right: ByVal(1),
					},
				},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
