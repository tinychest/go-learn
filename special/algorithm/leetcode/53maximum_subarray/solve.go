package _3maximum_subarray

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
// A subarray is a contiguous part of an array.
//
// Constraints:
//
//    1 <= nums.length <= 105
//    -104 <= nums[i] <= 104
//
//
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	// 并没有什么另辟蹊径的想法，但是有联想到 11
	// - 这注定至少需要完整遍历一遍数组的
	// 分而治之先不考虑，看看能否想出 O(n) 的做法

	l := len(nums)

	var totalMax, total int
	for i := 0; i < l; i++ {
		total += nums[i]
		if totalMax < total {
			totalMax = total
		}
		// 没有下边这个，就得到了端点在最左边的最长结果
		// 如何得到任意节点开始的，满足题意的结果呢
		// 只要发现某一段成为了累赘，直接丢弃解可，加一个 total 的变更处理就行了
		if total < 0 {
			total = 0
		}
	}

	return totalMax
}

// 简单调整了一下就符合题目需求了
// Runtime 190 ms
// Memory 10.1 MB
// - 算法的时间复杂度和空间复杂度都是达标的
// - Forward 就不做了，做的话肯定还可以进一步节约时间
func maxSubArray2(nums []int) int {
	l, total, totalMax := len(nums), 0, nums[0]

	for i := 0; i < l; i++ {
		if total < 0 {
			total = 0
		}
		total += nums[i]

		if totalMax < total {
			totalMax = total
		}
	}

	return totalMax
}
