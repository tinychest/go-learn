package sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
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
			InsertSort2(tt.args.arr)
			got := tt.args.arr

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
