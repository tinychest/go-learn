package _reverse_integer

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	t.Log(2 << 30)
	t.Log(int(math.Pow(2, 31)))
	t.Log(2<<30 - 1)
	t.Log(int(math.Pow(2, 31)) - 1)

	t.Log(-2 << 30)
	t.Log(-int(math.Pow(2, 31)))
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{0},
			want: 0,
		},
		{
			name: "test02",
			args: args{120},
			want: 21,
		},
		{
			name: "test03",
			args: args{-120},
			want: -21,
		},
		{
			name: "test04",
			args: args{-123},
			want: -321,
		},
		{
			name: "test05",
			args: args{123},
			want: 321,
		},
		{
			name: "test06",
			args: args{1534236469},
			want: 0,
		},
		{
			name: "test07",
			args: args{1563847412},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
