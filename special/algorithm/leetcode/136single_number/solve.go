package _36single_number

import "sort"

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//
// You must implement a solution with a linear runtime complexity and use only constant extra space.
//
//
// Constraints:
//
//    1 <= nums.length <= 3 * 104
//    -3 * 104 <= nums[i] <= 3 * 104
//    Each element in the array appears twice except for one element which appears only once.

// 就用 map，不被 use only constant extra space 限制
// 时间复杂度：O(1.5n)
// [结果]
// Runtime: 23 ms, faster than 66.86% of Go online submissions for Single Number.
// Memory Usage: 6.7 MB, less than 35.68% of Go online submissions for Single Number.
func singleNumber(nums []int) int {
	// return step1(nums)
	return step2(nums)
}

func step1(nums []int) int {
	m := make(map[int]bool, len(nums)/2+1)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = false
		} else {
			m[v] = true
		}
	}
	for k, v := range m {
		if !v {
			return k
		}
	}
	// never happened
	return -1
}

// 使用经典的排序来做，时间复杂度 O(log2_n + n)
func step2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

// 看过参考答案后，觉得最给力的答案是这个
// - 异或：相同的数异或得到 0，0 和任何数异或得到任何数（真棒）
// - 还有一个打小广告的莫名奇妙的广告
func refer1(nums []int) int {
	n := nums[0]

	for i := 1; i < len(nums); i++ {
		n ^= nums[i]
	}

	return n
}

