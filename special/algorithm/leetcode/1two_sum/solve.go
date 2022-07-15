package _two_sum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
//
// Constraints:
//
//    2 <= nums.length <= 104
//    -109 <= nums[i] <= 109
//    -109 <= target <= 109
//    Only one valid answer exists.
//
//
// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		// 这个一定要放在后边，不然就会出现元素重复使用的情况
		m[v] = i
	}
	panic("404")
}
