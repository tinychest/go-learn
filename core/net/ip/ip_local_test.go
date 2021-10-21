package ip

import "testing"

func TestIsLocal(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{
				ipStr: "localhost",
			},
			want: false,
		},
		{
			name: "test02",
			args: args{
				ipStr: "127.0.0.1",
			},
			want: true,
		},
		{
			name: "test03",
			args: args{
				ipStr: "::1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLoopback(tt.args.ipStr); got != tt.want {
				t.Errorf("IsLoopback() = %v, want %v", got, tt.want)
			}
		})
	}
}
