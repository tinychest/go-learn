package _6permutations

import (
	"reflect"
	"testing"
)

// 题目没有要求结果的顺序，所以尽管结果不是完全匹配，但下面的样例是符合题目要求的
func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case01",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "case02",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "case03",
			args: args{
				nums: []int{1},
			},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
