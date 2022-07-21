package _5jump_game_ii

import "testing"

func Test_jump(t *testing.T) {
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
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "case02",
			args: args{
				nums: []int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
		{
			name: "fail01",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "fail02",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "fail03",
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
