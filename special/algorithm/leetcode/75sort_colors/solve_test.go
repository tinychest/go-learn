package _5sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case01",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "case02",
			args: args{
				nums: []int{2, 0, 1},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "case03",
			args: args{
				nums: []int{1, 0},
			},
			want: []int{0, 1},
		},
		{
			name: "case04",
			args: args{
				nums: []int{1, 1, 1, 1, 0},
			},
			want: []int{0, 1, 1, 1, 1},
		},
		{
			name: "case05",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sortColors(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("minDistance() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
