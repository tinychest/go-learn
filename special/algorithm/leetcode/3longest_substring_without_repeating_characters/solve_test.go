package _longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "case02",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "case03",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "case04",
			args: args{s: "aabaab!bb"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
