package decimal

import (
	"testing"
)

func TestBadCeil(t *testing.T) {
	type args struct {
		s float64
		b int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "01", args: args{0.0700_01, 4}, want: 0.0701},
		{name: "02", args: args{0.0700_1, 4}, want: 0.0701},
		{name: "Fail01", args: args{0.07, 2}, want: 0.07},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BadCeil(tt.args.s, tt.args.b); got != tt.want {
				t.Errorf("BadCeil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBadRound(t *testing.T) {
	type args struct {
		s float64
		b int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "01", args: args{0.0700_49, 4}, want: 0.07},
		{name: "02", args: args{0.0700_5, 4}, want: 0.0701},
		{name: "Fail01", args: args{float64(0.03) + 0.005, 2}, want: 0.04},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BadRound(tt.args.s, tt.args.b); got != tt.want {
				t.Errorf("BadRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBadFloor(t *testing.T) {
	type args struct {
		s float64
		b int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "01", args: args{0.07, 4}, want: 0.07},
		{name: "02", args: args{0.0700_01, 4}, want: 0.07},
		{name: "03", args: args{0.0700_99, 4}, want: 0.07},
		{name: "Fail01", args: args{float64(0.1) + float64(0.7), 1}, want: 0.8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BadFloor(tt.args.s, tt.args.b); got != tt.want {
				t.Errorf("BadFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}
