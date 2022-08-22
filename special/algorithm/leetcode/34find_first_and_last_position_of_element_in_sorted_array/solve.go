package _4find_first_and_last_position_of_element_in_sorted_array

// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
//
// If target is not found in the array, return [-1, -1].
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Constraints:
//
//    0 <= nums.length <= 105
//    -109 <= nums[i] <= 109
//    nums is a non-decreasing array.
//    -109 <= target <= 109

func searchRange(nums []int, target int) []int {
	// return step01(nums, target)
	return step02(nums, target)
}

// 思路：二分查找到目标数后，向前向后查找下标
// 但好像时间复杂度是不可预测的，即不符合题目需求
//
// 完成下面的思路后，收获：
// - 该题很经典，当二分找的不仅是单个值得时候，应该从中吸收沉淀
// - 如果题目要求以指定时间复杂度去实现，除了完全符合，还可以增加常量次数的同时间复杂度（应试着向着该方向去思考）
func step01(nums []int, target int) []int {
	idx := halfSearch(nums, target)
	if idx == -1 {
		return []int{idx, idx}
	}

	start, end := idx, idx
	for i := idx; i >= 0; i-- {
		if nums[i] != target {
			break
		}
		start = i
	}
	for i := idx; i < len(nums); i++ {
		if nums[i] != target {
			break
		}
		end = i
	}
	return []int{start, end}
}

// 寻找前后的边界也应该用折半的思想去找，使用测试用例完善下面的代码，花了不少时间
// 但是提交效果确实可以
func step02(nums []int, target int) []int {
	idx := halfSearch(nums, target)
	if idx == -1 {
		return []int{idx, idx}
	}

	dstStart, dstEnd := idx, idx

	// 在前面，寻找第一个出现的 target
	start, mid, end := 0, idx/2, idx
	for {
		if nums[mid] != target && nums[mid+1] == target {
			dstStart = mid + 1
			break
		}
		if nums[mid] == target && (mid == start || nums[mid-1] != target) {
			dstStart = mid
			break
		}
		if nums[mid] != target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (start + end) / 2
	}

	// 在后面，寻找最后一个出现的 target
	start, end = idx, len(nums)-1
	mid = (start + end) / 2
	for {
		if nums[mid] != target && nums[mid-1] == target {
			dstEnd = mid - 1
			break
		}
		if nums[mid] == target && (mid == end || nums[mid+1] != target) {
			dstEnd = mid
			break
		}
		if nums[mid] != target {
			end = mid - 1
		} else {
			start = mid + 1
		}
		mid = (start + end) / 2
	}

	return []int{dstStart, dstEnd}
}

// O(log n) 的参考解法
// - 折半查找两次
// - 当与目标值相同时，不停止，而是向着既定的方向继续找
func refer(nums []int, target int) []int {
	// 等待自己实现
	return []int{}
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
