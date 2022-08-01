package _8subsets

// Given an integer array nums of unique elements, return all possible subsets (the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
//
// Constraints:
//
//    1 <= nums.length <= 10
//    -10 <= nums[i] <= 10
//    All the numbers of nums are unique.

// 和 22 题是类似的题，但其实更简单一些，但是好像提交结果并不理想
// 参考，还是关键利用递归的方式实现，dfs（Depth-First Search） 的思想
func subsets(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {
		for _, s := range res {
			t := make([]int, len(s)+1)
			copy(t, s)
			t[len(s)] = nums[i]
			// append 并没有影响 for range
			res = append(res, t)
		}
	}
	return res
}
