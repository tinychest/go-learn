package _5jump_game

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
//
//
// Constraints:
//
//    1 <= nums.length <= 104
//    0 <= nums[i] <= 105

// 题目顺序也是奇怪，这个题的增强版是 45 题
// 那现在做，不可得秒杀
//
// [结果]
// - 定义 max 函数
// Runtime: 85 ms, faster than 68.77% of Go online submissions for Jump Game.
// Memory Usage: 7.6 MB, less than 53.95% of Go online submissions for Jump Game.
// - 直接内联
// Runtime: 84 ms, faster than 69.77% of Go online submissions for Jump Game.
// Memory Usage: 7.5 MB, less than 61.86% of Go online submissions for Jump Game.
//
// 不是很理解为什么打不过用时更短的做法
func canJump(nums []int) bool {
	n := len(nums)

	// 能到达的最远下标
	furthest := 0
	for i := 0; i <= furthest; i++ {
		if furthest >= n-1 {
			return true
		}
		if furthest < i+nums[i] {
			furthest = i + nums[i]
		}
	}
	return false
}
