package _0regular_expression_matching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test00",
			args: args{
				s: "",
				p: ".*",
			},
			want: true,
		},
		{
			name: "test01",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "test02",
			args: args{
				s: "a",
				p: "aa",
			},
			want: false,
		},
		{
			name: "test03",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "test04",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "test05",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "test06",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "test07",
			args: args{
				s: "aaa",
				p: "a*a",
			},
			want: true,
		},
		{
			name: "test08",
			args: args{
				s: "a",
				p: "a*a*a",
			},
			want: true,
		},
		{
			name: "test09",
			args: args{
				s: "aaa",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "test10",
			args: args{
				s: "a",
				p: ".*..a*",
			},
			want: false,
		},
		{
			name: "test11",
			args: args{
				s: "",
				p: ".*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatchr() = %v, want %v", got, tt.want)
			}
		})
	}
}
