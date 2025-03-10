package _9divide_two_integers

import (
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{dividend: 10, divisor: 3},
			want: 3,
		},
		{
			name: "case02",
			args: args{dividend: 7, divisor: -3},
			want: -2,
		},
		{
			name: "case03",
			args: args{dividend: -8, divisor: -3},
			want: 2,
		},
		{
			name: "case04",
			args: args{dividend: -int(math.Pow(2, 31)), divisor: int(math.Pow(2, 31) - 1)},
			want: -1,
		},
		{
			name: "fail01",
			args: args{dividend: 1, divisor: 1},
			want: 1,
		},
		{
			name: "fail02",
			args: args{dividend: -2147483648, divisor: -1},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
