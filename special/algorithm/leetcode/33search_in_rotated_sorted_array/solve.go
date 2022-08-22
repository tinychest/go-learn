package _3search_in_rotated_sorted_array

// There is an integer array nums sorted in ascending order (with distinct values).
//
// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
//
// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Constraints:
//
//    1 <= nums.length <= 5000
//    -104 <= nums[i] <= 104
//    All values of nums are unique.
//    nums is an ascending array that is possibly rotated.
//    -104 <= target <= 104

// 题目的意思是，在一个经过旋转的有序数据组查找指定的元素值是否存在
func search(nums []int, target int) int {
	// return step01(nums, target)
	return refer(nums, target)
}

// 题目的期望是 O(log n)，看来肯定有巧妙的思想直接基于 rotate 的数组进行操作
// 在一个普通的有序数组中查找数，最佳的算法是折半查找
// - 思路1：找到旋转的目标位置，将其逆旋转得到一个普通的有序数组，然后使用折半查找
// - 思路2：找到旋转的目标位置，对前段使用折半，对后段使用折半
func step01(nums []int, target int) int {
	frontEnd := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			frontEnd = i + 1
		}
	}

	if res := halfSearch(nums[:frontEnd], target); res != -1 {
		return res
	}
	if res := halfSearch(nums[frontEnd:], target); res != -1 {
		return frontEnd + res
	}
	return -1
}

// 上面的做法本质就是这个，完全谈不上高效
func step02(nums []int, target int) int {
	frontEnd := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > nums[i+1] {
			frontEnd = i + 1
		}
	}

	if res := halfSearch(nums[frontEnd:], target); res != -1 {
		return frontEnd + res
	}
	return -1
}

// 时间复杂度为 O(log n) 的参考答案，稍微有点难理解
func refer(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		// 旋转了超过一半 或 转了一圈（没有旋转）
		// 特点：前半段 MIN：nums[start]、MAX：nums[mid]
		} else if nums[mid] >= nums[start] {
			// 目标数比中位数小，且比开头小，即比前半部分的数都小
			if target < nums[mid] && nums[start] <= target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		// 旋转了一点点（没有超过一半）
		// 特点：后半段 MIN：nums[mid]、MAX：nums[end]
		} else {
			// 目标数比中位数大，且比结尾大，即比后半部分的数都大
			if target > nums[mid] && nums[end] >= target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func halfSearch(nums []int, target int) int {
	start, mid, end := 0, len(nums)/2, len(nums)-1

	for start <= end {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		}
		if nums[mid] < target {
			start = mid + 1
		}
		mid = (start + end) / 2
	}
	return -1
}
