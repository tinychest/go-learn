package other

import (
	"reflect"
	"testing"
)

func Test_nextArr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case01",
			args: args{
				s: "abc",
			},
			want: []int{-1, -1, -1},
		},
		{
			name: "case02",
			args: args{
				s: "aaaab",
			},
			want: []int{-1, 0, 1, 2, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextArr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
