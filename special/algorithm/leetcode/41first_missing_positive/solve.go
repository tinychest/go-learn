package _1first_missing_positive

// Given an unsorted integer array nums, return the smallest missing positive integer.
//
// You must implement an algorithm that runs in O(n) time and uses constant extra space.
//
//
// Constraints:
//
//    1 <= nums.length <= 105
//    -2^31 <= nums[i] <= 2^31 - 1

// [思路]
// - 希望遍历一遍就想拿到结果，那肯定要将遍历过的数记录起来
//   但是只能借助有限的空间，那就用位来模拟了，而即使用位模拟，使用的空间也超标了
// - 或者说也可以借助特殊的位运算，一下得到结果，好像没那么简单
//
// [参考]
// 没办法，上 B 站搜了一下，既然空间有限，那就直接使用原数组空间
//   将值在数组长度范围内的数字，移动到匹配的位置上，之后再遍历这个数组，如果值不等于下标值，说明是缺失的值
//
// [小结]
// 不用特别关照 0，也不用特殊考虑重复的数值，注意题目的目的
// 有时间再来回顾一下，这就是 0(n) 时间复杂度找到第一个缺失的整数的解法
// 利用参数空间
func firstMissingPositive(nums []int) int {
	n := len(nums)
	var v int

	for i := 0; i < n; i++ {
		v = nums[i]
		// 不加最后一个判断条件，会因为 当前值和要交换的值相等 或者 当前值正好已经处于正确位置上了 导致死循环（容易忽略）
		if v < n && v > 0 && v != nums[v-1] {
			nums[i], nums[v-1] = nums[v-1], nums[i]
			// 必须要停留指针，不能会导致即将要处理的数值，被跳过去了
			i--
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
