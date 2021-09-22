package ip

import "testing"

func TestIsInternal(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bad",
			args: args{
				ipStr: "abc",
			},
			want: false,
		},
		{
			name: "test01",
			args: args{
				ipStr: "127.0.0.1",
			},
			want: true,
		},
		{
			name: "test02",
			args: args{
				ipStr: "192.168.1.125",
			},
			want: true,
		},
		{
			name: "test03",
			args: args{
				ipStr: "172.20.171.103",
			},
			want: true,
		},
		{
			name: "test04",
			args: args{
				ipStr: "58.33.192.66",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInternalIP(tt.args.ipStr); got != tt.want {
				t.Errorf("IsInternalIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
