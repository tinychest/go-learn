package _22coin_change

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "case02",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "case03",
			args: args{
				coins:  []int{1},
				amount: 0,
			},
			want: 0,
		},
		{
			name: "case04",
			args: args{
				coins:  []int{1},
				amount: 99,
			},
			want: 99,
		},
		{
			name: "case05",
			args: args{
				coins:  []int{8, 5, 4},
				amount: 9,
			},
			want: 2,
		},
		// 最优解是：5*419 8*408 3*186 4*83
		{
			name: "case06",
			args: args{
				coins:  []int{186, 419, 83, 408},
				amount: 6249,
			},
			want: 20,
		},
		// {
		// 	name: "case07",
		// 	args: args{
		// 		coins:  []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422},
		// 		amount: 9864,
		// 	},
		// 	want: ,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
