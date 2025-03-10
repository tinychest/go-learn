package _38counting_bits

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
// ans[i] is the number of 1's in the binary representation of i.
//
//
//
// Example 1:
//
// Input: n = 2
// Output: [0,1,1]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
//
// Example 2:
//
// Input: n = 5
// Output: [0,1,1,2,1,2]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101
//
// Constraints:
//
//    0 <= n <= 105
//
//
//
// Follow up:
//
//    It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in linear time O(n) and possibly in a single pass?
//    Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?

// 0 --> 0
// 1 --> 1
//
// 2 --> 10 1
// 3 --> 11 2
//
// 4 --> 100 1
// 5 --> 101 2
// 6 --> 110 2
// 7 --> 111 3
//
// 8 --> 1000 1
// 9 --> 1001 2
// 10 --> 1010 2
// 11 --> 1011 3
// 12 --> 1100 2
// 13 --> 1101 3
// 14 --> 1110 3
// 15 --> 1111 4
func countBits(n int) []int {
	res := make([]int, n+1, n+1)

	// 当前段的数字个数（2^0）
	duan := 1
	// 当前段的截至下标（0 + 2^0）
	duanIdx := 1
	// 当前段的中间下标（0 + 2^0 / 2）
	duanMidIdx := 0

	for i := 1; i <= n; i++ {
		if i <= duanMidIdx {
			res[i] = res[i-duan/2]
		} else {
			res[i] = 1 + res[i-duan/2]
		}

		if i == duanIdx {
			duanMidIdx = duanIdx + duan
			duan *= 2
			duanIdx += duan
		}
	}
	return res
}

// 上面的解法，是自己找的数值规律，这里才应该是最佳的，不解释，很好理解
func pref(n int) []int {
	res := make([]int, n+1, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i<<2] + i&1
	}
	return res
}
