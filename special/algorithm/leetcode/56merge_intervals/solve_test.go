package _6merge_intervals

import (
	"reflect"
	"testing"
)

func Test_mergeTwo(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case01",
			args: args{
				a: []int{1, 2},
				b: []int{3, 4},
			},
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "case02",
			args: args{
				a: []int{1, 1},
				b: []int{3, 4},
			},
			want: [][]int{{1, 1}, {3, 4}},
		},
		{
			name: "case04",
			args: args{
				a: []int{1, 3},
				b: []int{3, 4},
			},
			want: [][]int{{1, 4}},
		},
		{
			name: "case05",
			args: args{
				a: []int{1, 3},
				b: []int{1, 2},
			},
			want: [][]int{{1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwo(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case01",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "case02",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "case03",
			args: args{
				intervals: [][]int{{1, 1}},
			},
			want: [][]int{{1, 1}},
		},
		{
			name: "case04",
			args: args{
				intervals: [][]int{{1, 1}, {1, 2}},
			},
			want: [][]int{{1, 2}},
		},
		{
			name: "fail01",
			args: args{
				intervals: [][]int{{1, 4}, {0, 2}, {3, 5}},
			},
			want: [][]int{{0, 5}},
		},
		{
			name: "fail02",
			args: args{
				intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			},
			want: [][]int{{1, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
