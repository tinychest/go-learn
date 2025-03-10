package other

import (
	"reflect"
	"testing"
)

func TestStrMatch(t *testing.T) {
	type args struct {
		target string
		sub    string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case01",
			args: args{
				target: "abcdabcd",
				sub:    "abc",
			},
			want: []int{0, 4},
		},
		{
			name: "case02",
			args: args{
				target: "aabaac",
				sub:    "aaa",
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrMatch(tt.args.target, tt.args.sub); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
