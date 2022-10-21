package _69major_element

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
//
// Constraints:
//
//    n == nums.length
//    1 <= n <= 5 * 104
//    -109 <= nums[i] <= 109
//
//
// Follow-up: Could you solve the problem in linear time and in O(1) space?

// - 思路1：最无脑的就是用个 map 了，但是不符合进阶的要求
// - 分析：含有特定整数，其数量超过数组一半，如果对数组进行排序，那数组中间的元素肯定就是题目的答案了，但时间复杂度可就不会那么简单了
//   记得之前做过一道对所有元素进行异或进行解答的题，但是并不适用于这里
// - 参考：其实超过一半的重点是指定的整数其数量比剩余其他数的数量之和要多；细节是，出现 a b 到 b 时，当前记录的数应该是 b，且数量为 1
func majorityElement(nums []int) int {
	res, resSum := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if res != nums[i] {
			resSum--

			if resSum == 0 {
				res = nums[i]
				resSum = 1
			}
		} else {
			resSum++
		}
	}
	return res
}
