package _5sort_colors

// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
//
// You must solve this problem without using the library's sort function.
//
//
//
// Constraints:
//
//    n == nums.length
//    1 <= n <= 300
//    nums[i] is either 0, 1, or 2.
//
//
//
// Follow up: Could you come up with a one-pass algorithm using only constant extra space?
func sortColors(nums []int) {
	// step1(nums)
	// step2(nums)
	step3(nums)
}

// 没有去提交，因为时间复杂度是 O(n)
func step1(nums []int) {
	zero, one, two := 0, 0, 0
	for _, v := range nums {
		switch v {
		case 0:
			zero++
		case 1:
			one++
		case 2:
			two++
		}
	}
	nums = nums[0:0]
	for i := 0; i < zero; i++ {
		nums = append(nums, 0)
	}
	for i := 0; i < one; i++ {
		nums = append(nums, 1)
	}
	for i := 0; i < two; i++ {
		nums = append(nums, 2)
	}
}

// 0 往前边丢，2 往后边丢
// Runtime: 2 ms, faster than 46.98% of Go online submissions for Sort Colors.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Sort Colors.
func step2(nums []int) {
	start, end := 0, len(nums)-1

	for i := 0; i <= end; i++ {
		if nums[i] == 0 {
			if start != i {
				nums[start], nums[i] = nums[i], nums[start]
			}
			start++
		} else if nums[i] == 2 {
			nums[end], nums[i] = nums[i], nums[end]
			end--
			i--
		}
	}
}

// 参考了时间最短的算法的启示
// Runtime: 3 ms, faster than 34.27% of Go online submissions for Sort Colors.
// Memory Usage: 2.1 MB, less than 71.57% of Go online submissions for Sort Colors.
// 额，性能还下降了
// 时间和空间复杂度都达标，打住
func step3(nums []int) {
	start, end := 0, len(nums)-1

	for i := 0; i <= end; i++ {
		if nums[i] == 0 {
			if start != i {
				nums[i] = nums[start]
				nums[start] = 0
			}
			start++
		} else if nums[i] == 2 {
			nums[i] = nums[end]
			nums[end] = 2
			end--
			i--
		}
	}
}
