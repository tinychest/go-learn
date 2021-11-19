package _midian_sorted_arrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "测试一",
			args: args{
				nums1: []int{1},
				nums2: []int{},
			},
			want: 1.0,
		},
		{
			name: "测试二",
			args: args{
				nums1: []int{2},
				nums2: []int{2},
			},
			want: 2.0,
		},
		{
			name: "测试三",
			args: args{
				nums1: []int{1, 3, 5},
				nums2: []int{},
			},
			want: 3.0,
		},
		{
			name: "测试四",
			args: args{
				nums1: []int{1, 3, 5},
				nums2: []int{1, 3, 5},
			},
			want: 3.0,
		},
		{
			name: "测试五",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.0,
		},
		{
			name: "测试五",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
