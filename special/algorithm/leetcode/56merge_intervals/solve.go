package _6merge_intervals

import (
	"sort"
)

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
//
// Constraints:
//
//    1 <= intervals.length <= 104
//    intervals[i].length == 2
//    0 <= starti <= endi <= 104

func merge(intervals [][]int) [][]int {
	// return step1(intervals)
	return step2(intervals)
}

// 以 区间的左端按照从小到大进行排序 为前提重新写一下这道题
// 会发现规律很好找，这样的实现效率就很高
func step2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	left, right := intervals[0][0], intervals[0][1]

	res := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if right >= intervals[i][0] {
			right = max(right, intervals[i][1])
		} else {
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	return append(res, []int{left, right})
}

// [思路]
// 将既定的区间放到结果集中，然后，向结果集不断添加新的区间，添加的同时遍历其他结果，判断区间的影响
// 并对相关的区间进行 拓展 或 合并
// 一个超大的区间对于每个子区间的合并结果都是每个子区间，也就是结果重复；为了解决这个，本来说是调整规则，实际感觉区间需要进行排序就行了
//
// [结果]
// Runtime: 243 ms, faster than 5.03% of Go online submissions for Merge Intervals.
// Memory Usage: 7 MB, less than 68.82% of Go online submissions for Merge Intervals.
//
// 死办法去实现的效果太惨烈了
func step1(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		conflict := false
		for _, v2 := range res {
			two := mergeTwo(intervals[i], v2)
			if len(two) == 1 {
				v2[0] = two[0][0]
				v2[1] = two[0][1]
				conflict = true
			}
		}
		if !conflict {
			res = append(res, intervals[i])
		}
	}
	return res
}

// 将两个区间合并起来
func mergeTwo(a, b []int) [][]int {
	a1, a2 := a[0], a[1]
	b1, b2 := b[0], b[1]

	// 相交
	if a1 >= b1 && a1 <= b2 || b1 >= a1 && b1 <= a2 {
		return [][]int{{min(a1, b1), max(a2, b2)}}
	}
	// 不相交
	return [][]int{a, b}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
