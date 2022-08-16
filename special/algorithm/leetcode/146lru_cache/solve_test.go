package _46lru_cache

import (
	"encoding/json"
	"reflect"
	"testing"
)

func PtrInt(i int) *int {
	return &i
}

func step(args1 []string, args2 [][]int) []*int {
	l := len(args1)

	res := make([]*int, 0, l)
	var c LRUCache

	for i := 0; i < l; i++ {
		switch args1[i] {
		case "LRUCache":
			c = Constructor(args2[i][0])
			res = append(res, nil)
		case "put":
			c.Put(args2[i][0], args2[i][1])
			res = append(res, nil)
		case "get":
			res = append(res, PtrInt(c.Get(args2[i][0])))
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
				args1: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				args2: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			want: []*int{nil, nil, nil, PtrInt(1), nil, PtrInt(-1), nil, PtrInt(-1), PtrInt(3), PtrInt(4)},
		},
		{
			name: "fail01",
			args: args{
				args1: []string{"LRUCache", "put", "put", "get", "put", "put", "get"},
				args2: [][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}},
			},
			want: []*int{nil, nil, nil, PtrInt(2), nil, nil, PtrInt(-1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := step(tt.args.args1, tt.args.args2); !reflect.DeepEqual(got, tt.want) {

				g, _ := json.Marshal(got)
				w, _ := json.Marshal(tt.want)
				// t.Errorf("step() = %v, want %v", got, tt.want)
				t.Errorf("step() = %v, want %v", string(g), string(w))
			}
		})
	}
}
