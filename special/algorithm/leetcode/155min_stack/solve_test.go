package _55min_stack

import (
	"encoding/json"
	"testing"
)

func p(i int) *int {
	return &i
}

func step(args1 []string, args2 [][]int) []*int {
	p := func(i int) *int { return &i }

	l := len(args1)

	res := make([]*int, 0, l)
	var s MinStack

	for i := 0; i < l; i++ {
		switch args1[i] {
		case "MinStack":
			s = Constructor()
			res = append(res, nil)
		case "push":
			s.Push(args2[i][0])
			res = append(res, nil)
		case "top":
			res = append(res, p(s.Top()))
		case "pop":
			s.Pop()
			res = append(res, nil)
		case "getMin":
			res = append(res, p(s.GetMin()))
		}
	}
	return res
}

func Test_step(t *testing.T) {
	type args struct {
		args1 []string
		args2 [][]int
	}
	tests := []struct {
		name string
		args args
		want []*int
	}{
		{
			name: "case01",
			args: args{
				args1: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
				args2: [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			},
			want: []*int{nil, nil, nil, nil, p(-3), nil, p(0), p(-2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := step(tt.args.args1, tt.args.args2)

			gotBs, _ := json.Marshal(got)
			wantBs, _ := json.Marshal(tt.want)

			if string(gotBs) != string(wantBs) {
				t.Errorf("step() = %v, want %v", string(gotBs), string(wantBs))
			}
		})
	}
}
