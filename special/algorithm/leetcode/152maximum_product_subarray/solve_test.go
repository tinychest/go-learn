package _52maximum_product_subarray

import "testing"

func Test_maxProduct(t *testing.T) {
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
				nums: []int{2, 3, -2, 4},
			},
			want: 6,
		},
		{
			name: "case02",
			args: args{
				nums: []int{-2, 0, -1},
			},
			want: 0,
		},
		{
			name: "case03",
			args: args{
				nums: []int{1, 3, 3, -2, 0, 1, -2, -3, 2},
			},
			want: 12,
		},
		{
			name: "case04",
			args: args{
				nums: []int{-2},
			},
			want: -2,
		},
		{
			name: "case05",
			args: args{
				nums: []int{-2, 0, -1},
			},
			want: 0,
		},
		{
			name: "case06",
			args: args{
				nums: []int{-3, -1, -1},
			},
			want: 3,
		},
		{
			name: "case07",
			args: args{
				nums: []int{-2, 0},
			},
			want: 0,
		},
		{
			name: "case08",
			args: args{
				nums: []int{3, 0, 0, 2, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
