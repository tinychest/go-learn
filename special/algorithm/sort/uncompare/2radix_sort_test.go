package uncompare

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case01",
			args: args{
				arr: []int{1, 2, 45, 12, 11, 21, 0},
			},
			want: []int{0, 1, 2, 11, 12, 21, 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RadixSort(tt.args.arr)
			got := tt.args.arr

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RadixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
