package _1first_missing_positive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
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
				nums: []int{1, 2, 0},
			},
			want: 3,
		},
		{
			name: "case02",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		},
		{
			name: "case03",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		},
		{
			name: "case04",
			args: args{
				nums: []int{-1, 1, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "case05",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "case06",
			args: args{
				nums: []int{-1, 4, 2, 1, 9, 10},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
