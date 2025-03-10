package _61hamming_distance

// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//
// Given two integers x and y, return the Hamming distance between them.
//
//
//
// Example 1:
//
// Input: x = 1, y = 4
// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//       ↑   ↑
// The above arrows point to positions where the corresponding bits are different.
//
// Example 2:
//
// Input: x = 3, y = 1
// Output: 1
//
//
//
// Constraints:
//
//    0 <= x, y <= 231 - 1

// 题目很简单
func hammingDistance(x int, y int) int {
	res := 0

	part := x ^ y
	for part > 0 {
		if part&1 == 1 {
			res++
		}
		part = part >> 1
	}

	return res
}
