package _palindrome_number

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "test02",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "test03",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "test04",
			args: args{
				x: -101,
			},
			want: false,
		},
		{
			name: "test05",
			args: args{
				x: 11,
			},
			want: true,
		},
		{
			name: "test06",
			args: args{
				x: 1001,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aisPalindrome2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
