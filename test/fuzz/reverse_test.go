package fuzz

import (
	"testing"
	"unicode/utf8"
)

// https://go.dev/security/fuzz/
// https://segmentfault.com/a/1190000041467510

// 环境要求：GO 1.18
// 测试方法签名要求：func FuzzXxx(f *testing.F)
//
// 自动生成的的单元测试方法：Test_reverse
// 手动编写的模糊测试方法：Fuzz_reverse（没有安装 Go 1.18 绿色小箭头点击都不出来运行的提示）
// 暂时搁置，等待真实的业务场景
//
// 运行模糊测试的命令：go test -fuzz={FuzzTestName}

func Fuzz_reverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := reverse(orig)
		doubleRev := reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case01",
			args: args{s: "abc"},
			want: "cba",
		},
		{
			name: "case02",
		},
		{
			name: "case03",
			args: args{s: "112233"},
			want: "332211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
