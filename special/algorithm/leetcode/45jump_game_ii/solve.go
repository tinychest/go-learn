package _5jump_game_ii

// Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
//
// Each element in the array represents your maximum jump length at that position.
//
// Your goal is to reach the last index in the minimum number of jumps.
//
// You can assume that you can always reach the last index.
//
//
// Constraints:
//
//    1 <= nums.length <= 104
//    0 <= nums[i] <= 1000

func jump(nums []int) int {
	// return step(nums)
	return refer(nums)
}

// Runtime: 30 ms, faster than 52.76% of Go online submissions for Jump Game II.
// Memory Usage: 6.3 MB, less than 41.71% of Go online submissions for Jump Game II.
// 算法的时间复杂度本身就比较高，确实不是什么好的写法
// 这个算是借助额外空间，定位到目标位置的写法
// 因为没有注意约束的情况，错了好多次
func step(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	n := len(nums)
	record := make([]int, n)

	for i, v := range nums {
		if v == 0 {
			continue
		}
		if i+v >= n-1 {
			return record[i] + 1
		}

		for j := 1; j <= v; j++ {
			if record[i]+1 < record[i+j] || record[i+j] == 0 {
				record[i+j] = record[i] + 1
			}
		}
	}

	panic("never happened")
}

// 思考一下算法可以改进的地方，关键是是否存在规律，可以规避掉子循环，也就是不需要借助额外的空间
// 点：单程遍历可以知道能够达到目的地的起跳位置
// 参考答案后明白，思路应当是从当前已经达到的节点，能够再达到的最远节点
// 其实不难想到，画画图就能明白
func refer(nums []int) int {
	n := len(nums)
	left, right, farthest := 0, 0, 0
	steps := 0

	for farthest < n-1 {
		for i := left; i <= right; i++ {
			farthest = max(farthest, i+nums[i])
		}

		left, right = right+1, farthest
		steps++
	}
	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
