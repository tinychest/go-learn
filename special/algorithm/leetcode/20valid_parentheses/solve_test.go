package _0valid_parentheses

import "testing"

// 获取内存分配次数
func TestMemAlloc(t *testing.T) {
	s := "123456"

	var i byte
	t.Log("Allocs:", int(testing.AllocsPerRun(1, func() {
		for _, v := range []byte(s) {
			i = v
		}
	})))
	t.Log(i)
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case01",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "case02",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "case03",
			args: args{s: "(]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
