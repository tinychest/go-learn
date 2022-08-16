package _48sort_list

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case01",
			args: args{
				head: FromArr(4, 2, 1, 3),
			},
			want: FromArr(1, 2, 3, 4),
		},
		{
			name: "case02",
			args: args{
				head: FromArr(-1, 5, 3, 4, 0),
			},
			want: FromArr(-1, 0, 3, 4, 5),
		},
		{
			name: "step03",
			args: args{
				head: FromArr(),
			},
			want: FromArr(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
