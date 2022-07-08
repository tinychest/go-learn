package _6unique_binary_search_trees

import "testing"

func Test_numTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "case02",
			args: args{n: 3},
			want: 5,
		},
		{
			name: "case03",
			args: args{n: 4},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
