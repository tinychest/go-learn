package _6three_sum_closest

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},
		{
			name: "test02",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "test03",
			args: args{
				nums:   []int{1, 1, 1, 1},
				target: 0,
			},
			want: 3,
		},
		{
			name: "test04",
			args: args{
				nums:   []int{1, 1, 1, 0},
				target: 100,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
