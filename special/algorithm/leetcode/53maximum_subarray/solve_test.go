package _3maximum_subarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{
				nums: []int{-1, -1, -1},
			},
			want: -1,
		},
		{
			name: "case02",
			args: args{
				nums: []int{-1, 1, -1},
			},
			want: 1,
		},
		{
			name: "case03",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "error-case01",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "error-case02",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
