package _2generate_parentheses

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case01",
			args: args{
				n: 1,
			},
			want: []string{"()"},
		},
		{
			name: "case02",
			args: args{
				n: 2,
			},
			want: []string{"(())", "()()"},
		},
		{
			name: "case03",
			args: args{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
