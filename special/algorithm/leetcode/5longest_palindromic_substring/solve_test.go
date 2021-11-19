package _longest_palindromic_substring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test01",
			args: args{
				s: "abaad",
			},
			want: "aba",
		},
		{
			name: "test02",
			args: args{
				s: "ac",
			},
			want: "a",
		},
		{
			name: "test03",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "test04",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "test05",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
		{
			name: "test06",
			args: args{
				s: "aaaa",
			},
			want: "aaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
