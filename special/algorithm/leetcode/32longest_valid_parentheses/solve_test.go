package _2longest_valid_parentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "case02",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "case03",
			args: args{
				s: "((())))(()())",
			},
			want: 6,
		},
		{
			name: "fail01",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
		{
			name: "case04",
			args: args{
				s: "(()()",
			},
			want: 4,
		},
		{
			name: "case04",
			args: args{
				s: "()())",
			},
			want: 4,
		},
		{
			name: "fail02",
			args: args{
				s: "(()(((()",
			},
			want: 2,
		},
		{
			name: "fail03",
			args: args{
				s: "()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
